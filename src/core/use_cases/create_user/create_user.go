package createuser

import (
	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (useCase *CreateUserUseCase) Execute(name string, email string, username string, password string) (*entities.User, error) {
	var user *entities.User = entities.NewUser(name, email, username, password)

	err := useCase.validateUser(user)

	if err != nil {
		return nil, err
	}

	user, err = useCase.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (useCase *CreateUserUseCase) validateUser(user *entities.User) error {
	err := &MissingAttributesError{}

	if user.Name == "" {
		err.AppendMissingAttribute("Name")
	}

	if user.Email == "" {
		err.AppendMissingAttribute("Email")
	}

	if user.Username == "" {
		err.AppendMissingAttribute("Username")
	}

	if user.Password == "" {
		err.AppendMissingAttribute("Password")
	}

	if err.AnyMissingAttribute() {
		return err
	} else {
		return nil
	}
}
