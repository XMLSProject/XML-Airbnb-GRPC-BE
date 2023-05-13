package handler

import (
	"context"
	"fmt"
	"res_init/proto/reservation"
	"res_init/service"
)

func NewReservationHandler(service *service.ResService) *ReservationHandler {
	return &ReservationHandler{
		ResService: service,
	}
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
