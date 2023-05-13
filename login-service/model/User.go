package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name,omitempty"`
	Surname  string             `json:"surname,omitempty"`
	Username string             `json:"username,omitempty"`
	Password string             `json:"password,omitempty"`
	Email    string             `json:"email,omitempty"`
	Role     string             `json:"role,omitempty"`
}

func (User *User) BeforeCreate(*mongo.Client) error {
	User.ID = primitive.NewObjectID()
	return nil
}
