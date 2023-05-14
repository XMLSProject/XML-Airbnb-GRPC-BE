package handler

import (
	"context"
	"fmt"
	"res_init/model"
	"res_init/proto/reservation"
	"res_init/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewReservationHandler(service *service.ResService) *ReservationHandler {
	return &ReservationHandler{
		ResService: service,
	}
}
func parseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["role"].(string)
	if !ok {
		return ""
	}

	return sub
}

func userClaimmFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["username"].(string)
	if !ok {
		return ""
	}

	return sub
}

func checkRole(ctx context.Context) string {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ""
	}

	tokenInfo, _ := parseToken(token)

	role := userClaimFromToken(tokenInfo)

	fmt.Println("role is: " + role)
	return role
}
func checkUsername(ctx context.Context) string {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ""
	}

	tokenInfo, _ := parseToken(token)

	role := userClaimmFromToken(tokenInfo)

	fmt.Println("role is: " + role)
	return role
}

type ReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
	ResService *service.ResService
}

func (h ReservationHandler) GreetFromReservation(ctx context.Context, request *reservation.Request) (*reservation.Response, error) {
	return &reservation.Response{
		Greeting: fmt.Sprintf("Hihi from reservation %s!", request.Name),
	}, nil
}
func (h ReservationHandler) Reserve(ctx context.Context, request *reservation.RequestForReserve) (*reservation.ResponseForReserve, error) {
	role := checkRole(ctx)
	if role == "User" {
		fmt.Println(request)
		var Reservation = model.Reservation{}
		usr := checkUsername(ctx)
		Reservation.GuestUsername = usr
		Reservation.GuestNumber = int(request.GetReserve().GuestNumber)
		if request.GetReserve().Acception == "automatic" {
			Reservation.Accepted = "1"
		}
		if request.GetReserve().Acception == "handle" {
			Reservation.Accepted = "0"
		}
		Reservation.Accommodation = request.GetReserve().Accommodation
		layout := "2006-01-02T15:04:05Z"
		Reservation.FromDate, _ = time.Parse(layout, request.GetReserve().FromDate)
		Reservation.ToDate, _ = time.Parse(layout, request.GetReserve().ToDate)
		//objectId, _ := primitive.ObjectIDFromHex(accoId)
		fmt.Println(Reservation.FromDate)
		bul := h.ResService.FindOneByDate(Reservation.FromDate, Reservation.ToDate)
		//bult := h.ResService.FindOneByDateTwo(Reservation.FromDate, Reservation.ToDate)
		if bul {
			h.ResService.Create(&Reservation)
			return &reservation.ResponseForReserve{
				Reserve: fmt.Sprintf("Succesfully created! %s", request),
			}, nil
		}
		return &reservation.ResponseForReserve{
			Reserve: fmt.Sprintf("Reservation exist for that dates"),
		}, nil

	}
	return &reservation.ResponseForReserve{
		Reserve: "Error",
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")
}
func (h ReservationHandler) DeleteReservation(ctx context.Context, request *reservation.RequestDeleteReservation) (*reservation.ResponseDeleteReservation, error) {
	role := checkRole(ctx)
	if role == "User" || role == "Host" {
		acc, _ := h.ResService.FindOne(request.Delres)
		if time.Now().Before(acc.FromDate) {
			h.ResService.DeleteReservation(request.Delres)
			return &reservation.ResponseDeleteReservation{
				Delres: "Deleted",
			}, nil
		}
		return &reservation.ResponseDeleteReservation{
			Delres: "Deleted",
		}, status.Errorf(codes.Aborted, "Date is not before reservation date")

	}
	return &reservation.ResponseDeleteReservation{
		Delres: "Error",
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")

}
func (h ReservationHandler) AcceptReservation(ctx context.Context, request *reservation.DeleteRequest) (*reservation.DeleteResponse, error) {
	role := checkRole(ctx)
	if role == "Host" {
		fmt.Println(request.Dlt + " acce")
		h.ResService.AcceptReservation(request.Dlt)
		return &reservation.DeleteResponse{
			Dlt: "Updated",
		}, nil
	}
	return &reservation.DeleteResponse{
		Dlt: "Error",
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")
}
func (h ReservationHandler) CheckReservations(ctx context.Context, request *reservation.Request) (*reservation.Response, error) {
	fmt.Println(request.Name + " acce")
	checker, _ := h.ResService.GetAllAccommodationsByCreator(request.Name)
	if checker {
		return &reservation.Response{
			Greeting: "There is no reservations",
		}, nil
	}
	return &reservation.Response{
		Greeting: "There are reservations",
	}, nil
}

func (h ReservationHandler) GetAllReservations(ctx context.Context, request *reservation.AllReservationsRequest) (*reservation.AllReservationsResponse, error) {
	fmt.Println(request.Nothing)
	checker, _ := h.ResService.GetAllReservationsByAcc(request.Nothing)
	fmt.Println(len(checker))
	var allAccoInfo []*reservation.AllReservationInfo
	for _, acco := range checker {
		accoInfo := &reservation.AllReservationInfo{
			Id:            acco.ID.Hex(),
			FromDate:      acco.FromDate.String(),
			ToDate:        acco.ToDate.String(),
			GuestNumber:   int32(acco.GuestNumber),
			Accommodation: acco.Accommodation,
			Accepted:      acco.Accepted,
			Acception:     acco.Acception,
		}
		allAccoInfo = append(allAccoInfo, accoInfo)
	}

	response := &reservation.AllReservationsResponse{
		AllAcco: allAccoInfo,
	}

	return response, nil
}

func (h ReservationHandler) CheckReservationsByDates(ctx context.Context, request *reservation.CheckRequest) (*reservation.CheckResponse, error) {
	var accoId = request.GetCheckInfo().AccoId
	var dateFrom = request.GetCheckInfo().DateFrom
	var dateTo = request.GetCheckInfo().DateTo

	layout := "2006-01-02T15:04:05Z"
	dateFromDate, _ := time.Parse(layout, dateFrom)
	dateToDate, _ := time.Parse(layout, dateTo)

	checker, _ := h.ResService.CheckReservationsByDates(accoId, dateFromDate, dateToDate)
	return &reservation.CheckResponse{
		CheckRes: checker,
	}, nil
}
