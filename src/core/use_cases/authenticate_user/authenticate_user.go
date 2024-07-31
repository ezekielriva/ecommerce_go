package authenticateuser

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	"github.com/ezekielriva/ecommerce_go/src/core/repositories"
	"github.com/ezekielriva/ecommerce_go/src/utils"
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

	err = useCase.hashPassword(&cred)

	if err != nil {
		return nil, err
	}

	user, err := useCase.userRepo.Authenticate(cred)

	if err != nil {
		return nil, err
	}

	err = useCase.generateAuthToken(&cred)

	if err != nil {
		return nil, err
	}

	user.Credentials = &cred
	user, err = useCase.userRepo.Save(user)

	if err != nil {
		return nil, err
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

func (useCase *AuthenticateUserUseCase) hashPassword(cred *entities.UserCredentials) error {
	var err error

	cred.HashedPassword, err = utils.HashPassword(cred.Password)

	if err != nil {
		return errors.Join(errors.New("hash password"), err)
	}

	return nil
}

func (useCase *AuthenticateUserUseCase) generateAuthToken(cred *entities.UserCredentials) error {
	randomToken := make([]byte, 32)
	_, err := rand.Read(randomToken)

	if err != nil {
		return err
	}

	cred.AuthToken = base64.URLEncoding.EncodeToString(randomToken)
	cred.AuthTokenExp = time.Now().Add(time.Minute * 60)

	return nil
}
