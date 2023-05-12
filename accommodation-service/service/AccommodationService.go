package service

import (
	"accomm_module/model"
	repo "accomm_module/repository"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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
