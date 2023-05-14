package repo

import (
	"context"
	"fmt"
	"res_init/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *ResRepository) CreateReservation(Reservation *model.Reservation) error {
	Reservation.BeforeCreate(repo.DatabaseConnection)
	Reservation.Accepted = "0"
	_, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").InsertOne(context.TODO(), &Reservation)
	if err != nil {
		return err
	}
	fmt.Println("Sucessfully created request for reservation")
	return nil
}

func (repo *ResRepository) DeleteReservation(id string) error {
	fmt.Println(id + " evo ga id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").DeleteOne(
		context.TODO(),
		bson.M{"_id": objectId},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully deleted")
	return nil
}
func (repo *ResRepository) AcceptReservation(id string) error {
	fmt.Println(id + " evo ga id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").UpdateOne(
		context.TODO(),
		bson.M{"_id": objectId},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "accepted", Value: "1"},
		}}},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully updated")
	return nil
}

func (repo *ResRepository) GetAllReservationsByAccommodation(accoId string) (bool, error) {
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.M{"accommodation": accoId})
	if err != nil {
		return true, fmt.Errorf("failed to get reservations: %v", err)
	}
	defer cursor.Close(context.Background())

	var reservations []model.Reservation
	if err := cursor.All(context.Background(), &reservations); err != nil {
		return true, fmt.Errorf("failed to decode reservations: %v", err)
	}
	for _, reservation := range reservations {
		// Access individual reservation properties using the dot notation
		if reservation.Accepted == "1" {
			fmt.Println(reservation.ToDate.Date())
			if time.Now().After(reservation.ToDate) {
				return true, nil
			}
			return false, nil
		}

	}

	return true, nil
}
