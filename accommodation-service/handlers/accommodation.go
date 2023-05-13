package handler

import (
	"accomm_module/model"
	"accomm_module/proto/accommodation"
	"accomm_module/service"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

func NewAccommodationHandler(service *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		AccommodationService: service,
	}
}

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	AccommodationService *service.AccommodationService
}

func (h AccommodationHandler) GreetFromAccommodation(ctx context.Context, request *accommodation.Request) (*accommodation.Response, error) {

	fmt.Println("Uso u hendler")
	var Acco = model.Accommodation{}
	Acco.Name = "Joca"
	//err := h.UserService.Create(&Userr)
	//fmt.Println(err)
	//fmt.Println("Evo greske")
	return &accommodation.Response{
		Greeting: fmt.Sprintf("Hihi from accommodation!"),
	}, nil
}

func (h AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.CreateAccommodationRequest) (*accommodation.CreateAccommodationResponse, error) {
	var Accommodation = model.Accommodation{}
	Accommodation.Name = request.GetReg().Name
	Accommodation.Location = request.GetReg().Location
	Accommodation.Benefits = request.GetReg().Benefits
	Accommodation.Photos = request.GetReg().Photos
	Accommodation.MinGuests = int(request.GetReg().MinGuests)
	Accommodation.MaxGuests = int(request.GetReg().MaxGuests)
	//Accommodation.Creator = request.GetReg().Creator

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, _ := parseToken(token)

	username := userClaimFromToken(tokenInfo)

	fmt.Println("User id: " + username)

	Accommodation.Creator = username
	fmt.Println("Kreiram acco u handleru: " + Accommodation.Name)
	fmt.Println("Creator: " + username)

	h.AccommodationService.Create(&Accommodation)
	return &accommodation.CreateAccommodationResponse{
		Reg: &accommodation.Accommodation{},
	}, nil
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

	sub, ok := claims["username"].(string)
	if !ok {
		return ""
	}

	return sub
}

func (h AccommodationHandler) EditAccommodation(ctx context.Context, request *accommodation.EditAccoRequest) (*accommodation.EditAccoResponse, error) {
	var accoId = request.GetReg().AccoId
	var availableFrom = request.GetReg().AvailableFrom
	var availableTo = request.GetReg().AvailableTo
	var price = request.GetReg().Price
	var isPricePerGuest = request.GetReg().IsPricePerGuest

	layout := "2006-01-02T15:04:05Z"
	availableFromDate, _ := time.Parse(layout, availableFrom)
	objectId, _ := primitive.ObjectIDFromHex(accoId)
	availableToDate, _ := time.Parse(layout, availableTo)

	h.AccommodationService.EditPriceAndAvailability(objectId, availableFromDate, availableToDate, price, isPricePerGuest)
	return &accommodation.EditAccoResponse{
		Reg: &accommodation.EditAccoInfo{},
	}, nil
}

func (h AccommodationHandler) SearchAccommodation(ctx context.Context, request *accommodation.SearchAccoRequest) (*accommodation.SearchAccoResponse, error) {
	var location = request.GetSearchReqInfo().Location
	var dateFrom = request.GetSearchReqInfo().DateFrom
	var dateTo = request.GetSearchReqInfo().DateTo
	var guestNumber = request.GetSearchReqInfo().GuestNumber

	layout := "2006-01-02T15:04:05Z"
	dateFromDate, _ := time.Parse(layout, dateFrom)
	dateToDate, _ := time.Parse(layout, dateTo)

	accommodations, err := h.AccommodationService.SearchAccommodations(location, dateFromDate, dateToDate, int(guestNumber))
	if err != nil {
		fmt.Println("Error while searching accommodations")
		return nil, err
	}

	var searchInfo []*accommodation.SearchAccoInfo
	for _, acco := range accommodations {
		searchInfo = append(searchInfo, &accommodation.SearchAccoInfo{
			Id:              acco.ID.Hex(),
			Name:            acco.Name,
			Location:        acco.Location,
			Benefits:        acco.Benefits,
			Photos:          acco.Photos,
			MinGuests:       int32(acco.MinGuests),
			MaxGuests:       int32(acco.MaxGuests),
			AvailableFrom:   acco.AvailableFrom.Format(layout),
			AvailableTo:     acco.AvailableTo.Format(layout),
			Price:           acco.Price,
			IsPricePerGuest: acco.IsPricePerGuest,
			TotalPrice:      acco.TotalPrice,
		})
	}

	response := &accommodation.SearchAccoResponse{
		SearchInfo: searchInfo,
	}

	return response, nil
}
