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
	t.Run("Creates User with the provided data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mock_repositories.NewMockUserRepository(ctrl)
		useCase := NewCreateUserUseCase(repo)

		repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(user *entities.User) (any, error) {
			return user, nil
		}).Times(1)

		user, _ := useCase.Execute(
			"Name",
			"Email",
			"Username",
			"Password",
		)

		assert.Equal(t, "Name", user.Name, "Name doesnt match")
		assert.Equal(t, "Email", user.Email, "Email doesnt match")
		assert.Equal(t, "Username", user.Username, "Username doesnt match")
	})

	t.Run("Throw an Error when User Data is incomplete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		repo := mock_repositories.NewMockUserRepository(ctrl)
		useCase := NewCreateUserUseCase(repo)

		_, err := useCase.Execute("", "", "", "")

		assert.EqualError(t, err, errors.New("Missing attributes: Name, Email, Username, Password").Error())
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
