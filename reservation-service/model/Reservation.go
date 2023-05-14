package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reservation struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FromDate      time.Time          `bson:"fromdate,omitempty" json:"fromdate,omitempty"`
	ToDate        time.Time          `bson:"todate,omitempty" json:"todate,omitempty"`
	GuestNumber   int                `bson:"guestnumber,omitempty" json:"guestnumber,omitempty"`
	Accommodation string             `bson:"accommodation,omitempty" json:"accommodation,omitempty"`
	Accepted      string             `bson:"accepted,omitempty" json:"accepted,omitempty"`
	Acception     string             `bson:"acception,omitempty" json:"acception,omitempty"`
}

func (Reservation *Reservation) BeforeCreate(*mongo.Client) error {
	Reservation.ID = primitive.NewObjectID()
	return nil
}
