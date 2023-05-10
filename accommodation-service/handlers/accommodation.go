package handler

import (
	"accomm_module/model"
	"accomm_module/proto/accommodation"
	"accomm_module/service"
	"context"
	"fmt"
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
