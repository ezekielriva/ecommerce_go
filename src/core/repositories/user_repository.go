package repositories

import "github.com/ezekielriva/ecommerce_go/src/core/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
}
