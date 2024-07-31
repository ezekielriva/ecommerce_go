package createuser

import (
	"errors"
	"testing"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	mock_repositories "github.com/ezekielriva/ecommerce_go/src/core/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUseCase(t *testing.T) {
	validCases := []struct {
		desc     string
		name     string
		email    string
		username string
		password string
	}{
		{
			desc:     "Valid user with email and password",
			name:     "Name",
			email:    "Email",
			password: "Password",
		},
		{
			desc:     "Valid user with username and password",
			name:     "Name",
			username: "Username",
			password: "Password",
		},
	}

	for _, tC := range validCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repositories.NewMockUserRepository(ctrl)
			useCase := NewCreateUserUseCase(repo)

			repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(user *entities.User) (any, error) {
				user.Id = 1
				return user, nil
			}).Times(1)

			user, err := useCase.Execute(tC.name, tC.email, tC.username, tC.password)

			if err != nil {
				t.Error(err)
			}

			assert.NotEmpty(t, user, "Missing User")
			assert.Equal(t, tC.name, user.Name, "Name doesnt match")
			assert.Equal(t, tC.email, user.Credentials.Email, "Email doesnt match")
			assert.Equal(t, tC.username, user.Credentials.Username, "Username doesnt match")
			assert.NotEmpty(t, user.Credentials.HashedPassword, "No hashed password")
		})
	}

	t.Run("Throw an Error when User Data is incomplete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mock_repositories.NewMockUserRepository(ctrl)
		useCase := NewCreateUserUseCase(repo)

		_, err := useCase.Execute("", "", "", "")

		assert.EqualError(t, err, errors.New("Missing attributes: Name, Email or Username, Password").Error())
	})

	errorCases := []struct {
		desc string
		err  error
	}{{
		desc: "Throw an Error when User Already Exist",
		err:  errors.New("User already exist"),
	}}

	for _, tC := range errorCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repositories.NewMockUserRepository(ctrl)
			useCase := NewCreateUserUseCase(repo)

			repo.EXPECT().Create(gomock.Any()).Return(nil, tC.err).Times(1)

			_, err := useCase.Execute("Name", "Email", "Username", "Password")

			assert.EqualError(t, err, tC.err.Error())
		})
	}
}
