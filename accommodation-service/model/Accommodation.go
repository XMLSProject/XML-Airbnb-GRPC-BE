package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Accommodation struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Location  string    `json:"location,omitempty"`
	Benefits  string    `json:"benefits,omitempty"`
	Photos    []string  `json:"photos,omitempty"`
	MinGuests int       `json:"minGuests,omitempty"`
	MaxGuests int       `json:"maxGuests,omitempty"`
	Creator   string    `json:"creator,omitempty"`
}

func (Accommodation *Accommodation) BeforeCreate(*mongo.Client) error {
	Accommodation.ID = uuid.New()
	return nil
}
