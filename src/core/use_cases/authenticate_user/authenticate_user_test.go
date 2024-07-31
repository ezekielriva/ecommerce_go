package authenticateuser

import (
	"errors"
	"testing"
	"time"

	"github.com/ezekielriva/ecommerce_go/src/core/entities"
	mock_repositories "github.com/ezekielriva/ecommerce_go/src/core/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthenticateUserUseCase(t *testing.T) {
	password := "Password"

	validCases := []struct {
		desc     string
		creds    entities.UserCredentials
		useMocks bool
		err      error
		user     *entities.User
	}{
		{
			desc:     "When username and password match",
			creds:    entities.UserCredentials{Username: "Username", Password: password},
			user:     &entities.User{},
			useMocks: true,
		},
		{
			desc:     "When email and password match",
			creds:    entities.UserCredentials{Email: "Email", Password: password},
			user:     &entities.User{},
			useMocks: true,
		},
	}

	for _, tC := range validCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repositories.NewMockUserRepository(ctrl)
			useCase := NewAuthenticateUserUseCase(repo)

			if tC.useMocks {
				repo.EXPECT().Authenticate(gomock.Any()).DoAndReturn(func(_ entities.UserCredentials) (*entities.User, error) {
					return tC.user, tC.err
				})

				repo.EXPECT().Save(gomock.Any()).DoAndReturn(func(_ *entities.User) (*entities.User, error) {
					return tC.user, tC.err
				})
			}

			user, _ := useCase.Execute(tC.creds)

			assert.Equal(t, tC.user.Id, user.Id, "User IDs must match")
			assert.NotEmpty(t, user.Credentials.HashedPassword)
			assert.NotEmpty(t, user.Credentials.AuthToken)
			assert.GreaterOrEqual(t, user.Credentials.AuthTokenExp, time.Now())
			assert.LessOrEqual(t, user.Credentials.AuthTokenExp, time.Now().Add(time.Minute*60))
		})
	}

	errorCases := []struct {
		desc  string
		creds entities.UserCredentials
		setup func(repo *mock_repositories.MockUserRepository, user *entities.User, err error)
		err   error
		user  *entities.User
	}{
		{
			desc:  "When email and password doesnt match",
			creds: entities.UserCredentials{Email: "Email", Password: password, HashedPassword: gomock.Any().String()},
			err:   errors.New("User not found"),
			setup: func(repo *mock_repositories.MockUserRepository, user *entities.User, err error) {
				repo.EXPECT().Authenticate(gomock.Any()).DoAndReturn(func(_ entities.UserCredentials) (*entities.User, error) {
					return user, err
				})
			},
		},
		{
			desc:  "When username and password doesnt match",
			creds: entities.UserCredentials{Email: "Email", Password: password, HashedPassword: gomock.Any().String()},
			err:   errors.New("User not found"),
			setup: func(repo *mock_repositories.MockUserRepository, user *entities.User, err error) {
				repo.EXPECT().Authenticate(gomock.Any()).DoAndReturn(func(_ entities.UserCredentials) (*entities.User, error) {
					return user, err
				})
			},
		},
		{
			desc:  "When credentials are missing",
			creds: entities.UserCredentials{},
			err:   errors.New("validate user credentials params: email or username must be provided"),
			setup: func(repo *mock_repositories.MockUserRepository, user *entities.User, err error) {},
		},
		{
			desc:  "When password is missing",
			creds: entities.UserCredentials{Email: "Email"},
			err:   errors.New("validate user credentials params: password must be provided"),
			setup: func(repo *mock_repositories.MockUserRepository, user *entities.User, err error) {},
		},
	}

	for _, tC := range errorCases {
		t.Run(tC.desc, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_repositories.NewMockUserRepository(ctrl)
			useCase := NewAuthenticateUserUseCase(repo)

			tC.setup(repo, tC.user, tC.err)

			_, err := useCase.Execute(tC.creds)

			assert.Equal(t, err, tC.err, "Errors must match")
		})
	}
}
