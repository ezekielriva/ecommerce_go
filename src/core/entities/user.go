package entities

import "time"

type UserID int

type UserCredentials struct {
	Email          string
	Password       string
	Username       string
	HashedPassword string
	AuthToken      string
	AuthTokenExp   time.Time
}

type User struct {
	Id          UserID
	Name        string
	Credentials *UserCredentials
}

func NewUser(name string, email string, username string, password string) *User {
	return &User{
		Name: name,
		Credentials: &UserCredentials{
			Email:    email,
			Username: username,
			Password: password,
		},
	}
}
