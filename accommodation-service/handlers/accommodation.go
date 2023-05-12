package handler

import (
	"accomm_module/model"
	"accomm_module/proto/accommodation"
	"accomm_module/service"
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"

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
