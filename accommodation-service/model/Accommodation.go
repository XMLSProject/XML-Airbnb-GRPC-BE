package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Accommodation struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	Location        string             `bson:"location,omitempty" json:"location,omitempty"`
	Benefits        string             `bson:"benefits,omitempty" json:"benefits,omitempty"`
	Photos          []string           `bson:"photos,omitempty" json:"photos,omitempty"`
	MinGuests       int                `bson:"minGuests,omitempty" json:"minGuests,omitempty"`
	MaxGuests       int                `bson:"maxGuests,omitempty" json:"maxGuests,omitempty"`
	Creator         string             `bson:"creator,omitempty" json:"creator,omitempty"`
	AvailableFrom   time.Time          `bson:"availableFrom,omitempty" json:"availableFrom,omitempty"`
	AvailableTo     time.Time          `bson:"availableTo,omitempty" json:"availableTo,omitempty"`
	Price           float32            `bson:"price,omitempty" json:"price,omitempty"`
	IsPricePerGuest bool               `bson:"isPricePerGuest,omitempty" json:"isPricePerGuest,omitempty"`
	Acception       string             `bson:"acception,omitempty" json:"acception,omitempty"`
}

func (Accommodation *Accommodation) BeforeCreate(*mongo.Client) error {
	Accommodation.ID = primitive.NewObjectID()
	return nil
}
