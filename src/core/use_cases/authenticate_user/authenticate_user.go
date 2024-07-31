package authenticateuser

import (
	"errors"
	"fmt"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	authtokengeneration "github.com/ezekielriva/ecommerce_go/src/core/use_cases/auth_token_generation"
)

type AuthenticateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewAuthenticateUserUseCase(userRepo repositories.UserRepository) *AuthenticateUserUseCase {
	return &AuthenticateUserUseCase{
		userRepo: userRepo,
	}
}

/*
Authenticate Use Case:
 1. Validate Credentials
 2. Find User
 3. Generate Authentication Token
 4. Return User Data + Token Credentials
*/
func (useCase *AuthenticateUserUseCase) Execute(cred entities.UserCredentials) (*entities.User, error) {

	err := useCase.validateParams(cred)

	if err != nil {
		return nil, err
	}

	user, err := useCase.userRepo.Authenticate(cred)

	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	err = authtokengeneration.GenerateAuthToken(&cred)

	if err != nil {
		return nil, fmt.Errorf("generate authtoken: %w", err)
	}

	user.Credentials = &cred
	user, err = useCase.userRepo.Save(user)

	if err != nil {
		return nil, fmt.Errorf("credential update: %w", err)
	}

	return user, nil
}

func (useCase *AuthenticateUserUseCase) validateParams(cred entities.UserCredentials) error {
	if cred.Email == "" && cred.Username == "" {
		return errors.New("validate user credentials params: email or username must be provided")
	}

	if cred.Password == "" {
		return errors.New("validate user credentials params: password must be provided")
	}

	return nil
}
