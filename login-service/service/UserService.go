package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"first_init/model"
	repo "first_init/repository"
	"fmt"
	"regexp"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

func (service *UserService) FindUser(id string) (*model.User, error) {
	User, err := service.UserRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("user with id %s not found", id))
	}
	return &User, nil
}
func (service *UserService) FindUserForLogin(username string, password string) (*model.User, error) {
	var User *model.User
	User, err := service.UserRepo.FindUser(username, password)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("user with username and password not found"))
	}
	return User, nil
}
func registerValidation(User *model.User) bool {
	match1, _ := regexp.MatchString("[A-Z]{1}[a-z]+", User.Name)
	match2, _ := regexp.MatchString("[A-Z]{1}[a-z]+", User.Surname)
	match3, _ := regexp.MatchString("[A-Z]*[a-z]*[0-9]*", User.Username)
	match4, _ := regexp.MatchString("[A-Z]*[a-z]+[0-9]*@gmail.com", User.Email)

	return match1 && match2 && match3 && match4
}
func (service *UserService) checkUsername(User *model.User) bool {
	match := service.UserRepo.FindUsername(User.Username, User, User.Email)

	return match
}
func shaString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func (service *UserService) Create(User *model.User) error {
	regexError := registerValidation(User)
	User.Password = shaString(User.Password)
	if !regexError {
		fmt.Println("Regex error")
		return errors.New("regexError")
	}
	err := service.UserRepo.CreateUser(User)
	if err != nil {
		return err
	}
	return nil
}
func (service *UserService) FindUserByUsernameAndPassword(username string, password string) error {
	_, err := service.UserRepo.FindUser(username, password)
	if err != nil {
		return err
	}
	return nil
}
