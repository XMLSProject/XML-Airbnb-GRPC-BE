package repo

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"first_init/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DatabaseConnection *mongo.Client
}

func (repo *UserRepository) FindById(id string) (model.User, error) {
	User := model.User{}
	filter := bson.D{{Key: "name", Value: id}}
	fmt.Println("Doso do interakcije s bazom")
	err := repo.DatabaseConnection.Database("UserDB").Collection("users").FindOne(context.TODO(), filter).Decode(&User)
	fmt.Println("Proso interakcije s bazom")
	return User, err
}

func (repo *UserRepository) FindByUsername(username string) (model.User, error) {
	User := model.User{}
	filter := bson.D{{Key: "username", Value: username}}
	err := repo.DatabaseConnection.Database("UserDB").Collection("users").FindOne(context.TODO(), filter).Decode(&User)
	return User, err
}

func shaString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
func (repo *UserRepository) CreateUser(User *model.User) error {
	User.BeforeCreate(repo.DatabaseConnection)
	_, err := repo.DatabaseConnection.Database("UserDB").Collection("users").InsertOne(context.TODO(), &User)
	if err != nil {
		return err
	}
	// _, errr := repo.DatabaseConnection.Database("UserDB").Collection("users").Indexes().CreateOne(context.Background(),
	// 	mongo.IndexModel{
	// 		Keys:    bson.D{{Key: "email", Value: 1}},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// )
	// _, errrr := repo.DatabaseConnection.Database("UserDB").Collection("users").Indexes().CreateOne(context.Background(),
	// 	mongo.IndexModel{
	// 		Keys:    bson.D{{Key: "username", Value: 1}},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// )
	// if errr != nil {
	// 	return errors.New("email already exist")
	// }
	// if errrr != nil {
	// 	return errors.New("username already exist")
	// }
	fmt.Println("Sucessfully created")
	return nil
}
func (repo *UserRepository) UpdateUser(user *model.User) error {
	_, err := repo.DatabaseConnection.Database("UserDB").Collection("users").UpdateOne(
		context.TODO(),
		bson.M{"id": user.ID},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "name", Value: user.Name},
			{Key: "surname", Value: user.Surname},
			{Key: "username", Value: user.Username},
			{Key: "password", Value: user.Password},
		}}},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully updated")
	return nil
}
func (repo *UserRepository) DeleteUser(email string) error {
	_, err := repo.DatabaseConnection.Database("UserDB").Collection("users").DeleteOne(
		context.TODO(),
		bson.M{"username": email},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully deleted")
	return nil
}
func (repo *UserRepository) FindUsername(username string, User *model.User, email string) bool {
	filter := bson.D{{Key: "username", Value: username}}
	err := repo.DatabaseConnection.Database("UserDB").Collection("users").FindOne(context.TODO(), filter).Decode(&User)
	filterr := bson.D{{Key: "email", Value: email}}
	errr := repo.DatabaseConnection.Database("UserDB").Collection("users").FindOne(context.TODO(), filterr).Decode(&User)

	if err == nil {
		if errr == nil {
			return false
		}

	}
	return true
}
func (repo *UserRepository) FindUser(username string, password string) (*model.User, error) {
	var User model.User
	filter := bson.D{{Key: "username", Value: string(username)}, {Key: "password", Value: string(password)}}
	err := repo.DatabaseConnection.Database("UserDB").Collection("users").FindOne(context.TODO(), filter).Decode(&User)
	fmt.Println(User.Username + " " + User.Password)
	if err != nil {
		fmt.Println("User not found")
		return nil, errors.New("user not found")
	}
	if err == nil {
		fmt.Println("Sucessfully found")
	}
	return &User, err
}
