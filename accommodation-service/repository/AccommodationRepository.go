package repo

import (
	"accomm_module/model"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println("Pravim acco u repository: " + Accommodation.Name)
	_, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").InsertOne(context.TODO(), &Accommodation)
	if err != nil {
		fmt.Println("Greska u pravljenju u repository")
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Sucessfully created")
	return nil
}

func (repo *AccommodationRepository) EditPriceAndAvailability(accoId primitive.ObjectID, availableFrom time.Time, availableTo time.Time, price float32, isPricePerGuest bool) error {
	updateFields := bson.D{}

	if !availableFrom.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "availableFrom", Value: availableFrom})
	}

	if !availableTo.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "availableTo", Value: availableTo})
	}

	if price != 0 {
		updateFields = append(updateFields, bson.E{Key: "price", Value: price})
	}

	updateFields = append(updateFields, bson.E{Key: "isPricePerGuest", Value: isPricePerGuest})

	result, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").UpdateOne(
		context.TODO(),
		bson.M{"_id": accoId},
		bson.D{{Key: "$set", Value: updateFields}},
	)
	if err != nil {
		fmt.Println("Error editing flight with id: " + accoId.String())
		return err
	}
	if result.ModifiedCount == 0 {
		fmt.Println("No accommodation found with ID: " + accoId.String())
		return errors.New("no accommodation found")
	}

	fmt.Println("Successfully updated")
	return nil
}
