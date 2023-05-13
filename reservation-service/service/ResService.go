package service

import (
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
