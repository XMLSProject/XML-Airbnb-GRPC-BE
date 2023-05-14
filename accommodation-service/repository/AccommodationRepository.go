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

func (repo *AccommodationRepository) SearchAccommodations(location string, dateFrom time.Time, dateTo time.Time, guestNumber int) ([]model.SearchDTO, error) {
	filter := bson.D{
		{Key: "location", Value: location},
		{Key: "availableFrom", Value: bson.D{{Key: "$lte", Value: dateFrom}}},
		{Key: "availableTo", Value: bson.D{{Key: "$gte", Value: dateTo}}},
		{Key: "minGuests", Value: bson.D{{Key: "$lte", Value: guestNumber}}},
		{Key: "maxGuests", Value: bson.D{{Key: "$gte", Value: guestNumber}}},
	}

	cursor, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to search accommodations")
	}

	defer cursor.Close(context.Background())

	var accommodations []model.Accommodation
	if err = cursor.All(context.Background(), &accommodations); err != nil {
		fmt.Println("ERROR IN CURSOR: " + err.Error())
		return nil, fmt.Errorf("Failed to search accommodations")
	}
	var accoDTO []model.SearchDTO
	for _, acco := range accommodations {
		var totalPrice float32
		if acco.IsPricePerGuest {
			totalPrice = float32(guestNumber) * acco.Price * float32(dateTo.Sub(dateFrom).Hours()/24)
		} else {
			totalPrice = acco.Price * float32(dateTo.Sub(dateFrom).Hours()/24)
		}

		dto := model.SearchDTO{
			ID:              acco.ID,
			Name:            acco.Name,
			Location:        acco.Location,
			Benefits:        acco.Benefits,
			Photos:          acco.Photos,
			MinGuests:       acco.MinGuests,
			MaxGuests:       acco.MaxGuests,
			Creator:         acco.Creator,
			AvailableFrom:   acco.AvailableFrom,
			AvailableTo:     acco.AvailableTo,
			Price:           acco.Price,
			IsPricePerGuest: acco.IsPricePerGuest,
			TotalPrice:      totalPrice,
		}

		accoDTO = append(accoDTO, dto)
	}

	return accoDTO, nil
}

func (repo *AccommodationRepository) GetAllAccommodations() ([]model.Accommodation, error) {
	cursor, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Accommodation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return nil, fmt.Errorf("Failed to decode accommodations: %v", err)
	}

	return accommodations, nil
}
func (repo *AccommodationRepository) CheckHere(creator string) (string, error) {
	objectId, _ := primitive.ObjectIDFromHex(creator)
	cursor, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").Find(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return "", fmt.Errorf("failed to get accommodations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Accommodation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return "", fmt.Errorf("failed to decode accommodations: %v", err)
	}

	return accommodations[0].Acception, nil
}
func (repo *AccommodationRepository) GetAllAccommodationsByCreator(creator string) ([]model.Accommodation, error) {
	cursor, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").Find(context.Background(), bson.M{"creator": creator})
	if err != nil {
		return nil, fmt.Errorf("failed to get accommodations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Accommodation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return nil, fmt.Errorf("failed to decode accommodations: %v", err)
	}

	return accommodations, nil
}
func (repo *AccommodationRepository) DeleteAllAccom(creator string) error {
	_, err := repo.DatabaseConnection.Database("AccommodationsDB").Collection("accommodations").DeleteMany(context.Background(), bson.M{"creator": creator})
	if err != nil {
		return fmt.Errorf("failed to get accommodations: %v", err)
	}
	return nil
}
