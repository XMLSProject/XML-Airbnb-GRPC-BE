package repo

import (
	"context"
	"fmt"
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
