package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reservation struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name,omitempty"`
	Surname  string    `json:"surname,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Email    string    `json:"email,omitempty"`
	Role     string    `json:"role,omitempty"`
}

func (User *Reservation) BeforeCreate(*mongo.Client) error {
	User.ID = uuid.New()
	return nil
}
