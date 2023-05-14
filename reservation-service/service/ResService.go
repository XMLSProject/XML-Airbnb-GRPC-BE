package service

import (
	"fmt"
	"res_init/model"
	repo "res_init/repository"
	"time"
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
func (service *ResService) GetAllAccommodationsByCreator(creator string) (bool, error) {
	accommodations, err := service.ResRepo.GetAllReservationsByAccommodation(creator)
	if err != nil {
		return true, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}

func (service *ResService) GetAllReservationsByAcc(id string) ([]model.Reservation, error) {
	accommodations, err := service.ResRepo.GetAllReservationsByAcc(id)
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}

func (service *ResService) CheckReservationsByDates(accoId string, dateFrom time.Time, dateTo time.Time) (bool, error) {
	accommodations, err := service.ResRepo.CheckReservationsByDates(accoId, dateFrom, dateTo)
	if err != nil {
		return true, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}

func (service *ResService) FindOne(creator string) (*model.Reservation, error) {
	accommodations, err := service.ResRepo.FindOne(creator)
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}
func (service *ResService) FindOneByDate(fromDate time.Time, toDate time.Time) bool {
	accommodations := service.ResRepo.FindOneByDate(fromDate, toDate)

	return accommodations
}
