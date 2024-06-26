package entities

type UserID int

type User struct {
	Id       UserID
	Name     string
	Email    string
	Username string
	Password string
}

func NewUser(name string, email string, username string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
}
