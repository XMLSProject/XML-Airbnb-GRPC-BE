package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Accommodation struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name,omitempty"`
}

func (Accommodation *Accommodation) BeforeCreate(*mongo.Client) error {
	Accommodation.ID = uuid.New()
	return nil
}
