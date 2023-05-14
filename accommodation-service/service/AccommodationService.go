package service

import (
	"accomm_module/model"
	repo "accomm_module/repository"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	AccommodationRepo *repo.AccommodationRepository
}

func (service *AccommodationService) FindUser(id string) (*model.Accommodation, error) {
	fmt.Println("Uso u servis")
	Accommodation, err := service.AccommodationRepo.FindById(id)
	fmt.Println("Proso servis")
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("user with id %s not found", id))
	}
	return &Accommodation, nil
}

func shaString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func (service *AccommodationService) Create(Accommodation *model.Accommodation) error {
	fmt.Println("Pravim acco u servisu: " + Accommodation.Name)
	err := service.AccommodationRepo.CreateAccommodation(Accommodation)
	if err != nil {
		fmt.Println("GRESKA U PRAVLJENJU U SERVISU")
		return err
	}
	return nil
}

func (service *AccommodationService) EditPriceAndAvailability(accoId primitive.ObjectID, availableFrom time.Time, availableTo time.Time, price float32, isPricePerGuest bool) error {
	err := service.AccommodationRepo.EditPriceAndAvailability(accoId, availableFrom, availableTo, price, isPricePerGuest)

	if err != nil {
		fmt.Println("Error updating in service")
		return err
	}

	fmt.Println("Successfully updated")
	return nil
}

func (service *AccommodationService) SearchAccommodations(location string, dateFrom time.Time, dateTo time.Time, guestNumber int) ([]model.SearchDTO, error) {
	accommodations, err := service.AccommodationRepo.SearchAccommodations(location, dateFrom, dateTo, guestNumber)
	if err != nil {
		return nil, fmt.Errorf("Failed to search accommodations: %v", err)
	}

	return accommodations, nil
}

func (service *AccommodationService) GetAllAccommodations() ([]model.Accommodation, error) {
	accommodations, err := service.AccommodationRepo.GetAllAccommodations()
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}

func (service *AccommodationService) GetAllAccommodationsByCreator(creator string) ([]model.Accommodation, error) {
	accommodations, err := service.AccommodationRepo.GetAllAccommodationsByCreator(creator)
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}

	return accommodations, nil
}
func (repo *AccommodationService) CheckOne(id string) (string, error) {
	ret, err := repo.AccommodationRepo.CheckHere(id)
	fmt.Println("Proso interakcije s bazom")
	return ret, err
}
