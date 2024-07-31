package createuser

import (
	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	authtokengeneration "github.com/ezekielriva/ecommerce_go/src/core/use_cases/auth_token_generation"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

/*
Create User Use Case:
 1. Validate User Data
 2. Build Credentials Data
 3. Create User
 4. Return User
*/
func (useCase *CreateUserUseCase) Execute(name string, email string, username string, password string) (*entities.User, error) {
	user, err := entities.NewUser(name, email, username, password)

	if err != nil {
		return nil, err
	}

	err = useCase.validateUser(user)

	if err != nil {
		return nil, err
	}

	err = authtokengeneration.GenerateAuthToken(user.Credentials)

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
	err := useCase.validateUserAttributes(user)

	if err != nil {
		return err
	}

	return nil
}

func (userCase *CreateUserUseCase) validateUserAttributes(user *entities.User) error {
	err := &MissingAttributesError{}
	cred := user.Credentials

	if user.Name == "" {
		err.AppendMissingAttribute("Name")
	}

	if cred.Email == "" && cred.Username == "" {
		err.AppendMissingAttribute("Email or Username")
	}

	if cred.Password == "" {
		err.AppendMissingAttribute("Password")
	}

	if err.AnyMissingAttribute() {
		return err
	}

	return nil
}
