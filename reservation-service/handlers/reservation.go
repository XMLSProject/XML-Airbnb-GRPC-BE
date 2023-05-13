package handler

import (
	"context"
	"fmt"
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
func (h ReservationHandler) DeleteReservation(ctx context.Context, request *reservation.RequestDeleteReservation) (*reservation.ResponseDeleteReservation, error) {
	h.ResService.DeleteReservation(request.Delres)
	return &reservation.ResponseDeleteReservation{
		Delres: "Deleted",
	}, nil
}
