package service

import (
	"fmt"
	"res_init/model"
	repo "res_init/repository"
)

type ResService struct {
	ResRepo *repo.ResRepository
}

func (service *ResService) Create(Reservation *model.Reservation) error {
	err := service.ResRepo.CreateReservation(Reservation)
	if err != nil {
		return err
	}
	return nil
}
func (service *ResService) DeleteReservation(id string) error {
	err := service.ResRepo.DeleteReservation(id)

	if err != nil {
		return err
	}

	fmt.Println("Successfully deleted")
	return nil
}
func (service *ResService) AcceptReservation(id string) error {
	err := service.ResRepo.AcceptReservation(id)

	if err != nil {
		return err
	}

	fmt.Println("Successfully updated")
	return nil
}
