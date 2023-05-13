package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"res_init/model"
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
