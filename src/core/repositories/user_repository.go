package repositories

import "github.com/ezekielriva/ecommerce_go/src/core/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	Authenticate(creadentials entities.UserCredentials) (*entities.User, error)
	Save(user *entities.User) (*entities.User, error)
}
