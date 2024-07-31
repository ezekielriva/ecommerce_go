package entities

import (
	"time"

	"github.com/ezekielriva/ecommerce_go/src/utils"
)

type UserID int

type UserCredentials struct {
	Email          string
	Password       string
	Username       string
	HashedPassword string
	AuthToken      string
	AuthTokenExp   time.Time
}

func NewUserCredentials(email string, username string, password string) (*UserCredentials, error) {
	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	return &UserCredentials{
		Email:          email,
		Username:       username,
		Password:       password,
		HashedPassword: hashedPassword,
	}, nil
}

type User struct {
	Id          UserID
	Name        string
	Credentials *UserCredentials
}

func NewUser(name string, email string, username string, password string) (*User, error) {
	cred, err := NewUserCredentials(email, username, password)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:        name,
		Credentials: cred,
	}, nil
}
