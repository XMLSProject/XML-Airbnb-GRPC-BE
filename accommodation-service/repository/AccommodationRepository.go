package repo

import (
	"accomm_module/model"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccommodationRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *AccommodationRepository) FindById(id string) (model.Accommodation, error) {
	Accommodation := model.Accommodation{}
	filter := bson.D{{Key: "name", Value: id}}
	fmt.Println("Doso do interakcije s bazom")
	err := repo.DatabaseConnection.Database("AccommodationDB").Collection("accommodations").FindOne(context.TODO(), filter).Decode(&Accommodation)
	fmt.Println("Proso interakcije s bazom")
	return Accommodation, err
}

func shaString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func (repo *AccommodationRepository) CreateAccommodation(Accommodation *model.Accommodation) error {
	Accommodation.BeforeCreate(repo.DatabaseConnection)
	_, err := repo.DatabaseConnection.Database("AccommodationDB").Collection("accommodations").InsertOne(context.TODO(), &Accommodation)
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully created")
	return nil
}
