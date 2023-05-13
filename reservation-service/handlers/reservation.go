package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"res_init/model"
	"res_init/proto/reservation"
	"res_init/service"
	"time"
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
		Reservation.GuestNumber = int(request.GetReserve().GuestNumber)
		Reservation.Accepted = "0"
		Reservation.Accommodation = request.GetReserve().Accommodation
		layout := "2006-01-02T15:04:05Z"
		Reservation.FromDate, _ = time.Parse(layout, request.GetReserve().FromDate)
		Reservation.ToDate, _ = time.Parse(layout, request.GetReserve().ToDate)
		//objectId, _ := primitive.ObjectIDFromHex(accoId)
		h.ResService.Create(&Reservation)
		return &reservation.ResponseForReserve{
			Reserve: fmt.Sprintf("Succesfully created! %s", request),
		}, nil
	}
	return &reservation.ResponseForReserve{
		Reserve: "Error",
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")
}
func (h ReservationHandler) DeleteReservation(ctx context.Context, request *reservation.RequestDeleteReservation) (*reservation.ResponseDeleteReservation, error) {
	role := checkRole(ctx)
	if role == "User" || role == "Host" {
		h.ResService.DeleteReservation(request.Delres)
		return &reservation.ResponseDeleteReservation{
			Delres: "Deleted",
		}, nil
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
