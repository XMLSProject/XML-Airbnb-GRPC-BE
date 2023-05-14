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
	repo.CheckReservationsByDatesUpdate(id)

	fmt.Println("Successfully updated")
	return nil
}
func (repo *ResRepository) FindOne(creator string) (*model.Reservation, error) {
	objectId, _ := primitive.ObjectIDFromHex(creator)
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("failed to get accommodations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return nil, fmt.Errorf("failed to decode accommodations: %v", err)
	}

	return &accommodations[0], nil
}
func (repo *ResRepository) FindOneByDate(fromDate time.Time, toDate time.Time) bool {
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.D{})
	if err != nil {
		return true
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return true
	}
	for _, accommodation := range accommodations {
		// Perform the desired operation with each accommodation
		if (accommodation.FromDate.Before(fromDate) && accommodation.ToDate.After(fromDate)) || (accommodation.FromDate.Before(toDate) && accommodation.ToDate.After(toDate)) || (accommodation.FromDate.After(fromDate) && accommodation.FromDate.Before(toDate)) || (accommodation.ToDate.After(fromDate) && accommodation.ToDate.Before(toDate)) {
			return false
		}
	}

	return true
}
func (repo *ResRepository) FindOneByDateTwo(fromDate time.Time, toDate time.Time) bool {
	filter := bson.M{
		"fromDate": bson.M{"$lte": toDate},
		"toDate":   bson.M{"$gte": toDate},
	}
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), filter)
	if err != nil {
		return true
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return true
	}

	return false
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
func (repo *ResRepository) CheckReservationForUser(username string) (bool, error) {
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.M{"guestUsername": username})
	if err != nil {
		return true, fmt.Errorf("failed to get reservations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return true, fmt.Errorf("failed to decode accommodations: %v", err)
	}
	if len(accommodations) != 0 {
		return false, nil
	}

	return true, nil
}

func (repo *ResRepository) CheckReservationsByDates(accoId string, dateFrom time.Time, dateTo time.Time) (bool, error) {
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.M{"accommodation": accoId})
	if err != nil {
		return false, fmt.Errorf("failed to get reservations: %v", err)
	}
	defer cursor.Close(context.Background())
	var result = true
	var reservations []model.Reservation
	if err := cursor.All(context.Background(), &reservations); err != nil {
		return false, fmt.Errorf("failed to decode reservations: %v", err)
	}
	for _, reservation := range reservations {
		// Access individual reservation properties using the dot notation
		if reservation.Accepted == "1" {
			if (dateFrom.After(reservation.FromDate) && dateFrom.Before(reservation.ToDate)) || (dateTo.After(reservation.FromDate) && dateTo.Before(reservation.ToDate)) {
				result = false
			}
		}
	}
	return result, nil
}
func (repo *ResRepository) CheckReservationsByDatesUpdate(id string) bool {
	res, _ := repo.FindOne(id)
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.D{})
	if err != nil {
		return true
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return true
	}
	for _, accommodation := range accommodations {
		// Perform the desired operation with each accommodation
		if (accommodation.FromDate.Before(res.FromDate) && accommodation.ToDate.After(res.FromDate)) || (accommodation.FromDate.Before(res.ToDate) && accommodation.ToDate.After(res.ToDate)) || (accommodation.FromDate.After(res.FromDate) && accommodation.FromDate.Before(res.ToDate)) || (accommodation.ToDate.After(res.FromDate) && accommodation.ToDate.Before(res.ToDate)) {

			repo.DeleteReservation(res.ID.Hex())
			fmt.Println("obriso")
		}
	}

	return true
}
func (repo *ResRepository) GetAllReservationsByAcc(accomm string) ([]model.Reservation, error) {
	fmt.Println(accomm)
	cursor, err := repo.DatabaseConnection.Database("ReservationDB").Collection("reservations").Find(context.Background(), bson.M{"accommodation": accomm})
	if err != nil {
		return nil, fmt.Errorf("Failed to get accommodations: %v", err)
	}
	defer cursor.Close(context.Background())

	var accommodations []model.Reservation
	if err := cursor.All(context.Background(), &accommodations); err != nil {
		return nil, fmt.Errorf("Failed to decode accommodations: %v", err)
	}
	fmt.Println(len(accommodations))

	return accommodations, nil
}
