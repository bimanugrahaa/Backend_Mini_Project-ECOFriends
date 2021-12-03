package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"Backend_Mini_Project-ECOFriends/features/user/mocks"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userValue   user.UserCore
	usersValue  []user.UserCore
	userUseCase user.Bussiness
	userData    mocks.Data
)

func TestMain(m *testing.M) {
	userUseCase = NewUserBussiness(&userData)

	userValue = user.UserCore{
		ID:         1,
		Name:       "Barry Keoghan",
		Email:      "barrykeoghan1@gmail.com",
		Password:   "barrykeoghan1",
		Created_at: time.Now(),
	}

	usersValue = []user.UserCore{
		{
			ID:         1,
			Name:       "Barry Keoghan",
			Email:      "barrykeoghan1@gmail.com",
			Password:   "barrykeoghan1",
			Created_at: time.Now(),
		},
	}

	os.Exit(m.Run())
}

func TestGetAllUser(t *testing.T) {
	t.Run("valid - get all user", func(t *testing.T) {
		userData.On("SelectAllUser").Return(usersValue).Once()
		resp := userUseCase.GetAllUser()

		assert.NotEqual(t, len(resp), 0)
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("valid - get user by id", func(t *testing.T) {
		userData.On("SelectUserById", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		resp, _ := userUseCase.GetUserById(usersValue[0].ID)

		assert.Equal(t, resp.ID, usersValue[0].ID)
	})
}

func TestLogin(t *testing.T) {
	t.Run("valid - login", func(t *testing.T) {
		userData.On("Login", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		resp, err := userUseCase.Login(userValue)

		assert.Nil(t, err)
		assert.Equal(t, resp.Email, userValue.Email)
	})

	t.Run("invalid email - login", func(t *testing.T) {
		userData.On("Login", mock.AnythingOfType("user.UserCore")).Return(userValue, errors.New("something went wrong")).Once()
		_, err := userUseCase.Login(user.UserCore{
			Email:    "barrykeoghan2@gmail.com",
			Password: "barrykeoghan2",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "something went wrong")

	})
}

func TestCreateUser(t *testing.T) {
	t.Run("valid - create user", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		userData.On("InsertUser", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		_, err := userUseCase.CreateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.Nil(t, err)
	})

	t.Run("invalid - error create user", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		userData.On("InsertUser", mock.AnythingOfType("user.UserCore")).Return(userValue, errors.New("error insert user")).Once()
		_, err := userUseCase.CreateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error insert user")
	})

	t.Run("invalid email - data is available", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		_, err := userUseCase.CreateUser(userValue)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data is available")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("valid - update user", func(t *testing.T) {
		userData.On("SelectUserById", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		userData.On("EditUser", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		resp, err := userUseCase.UpdateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Email, userValue.Email)
	})

	t.Run("invalid - error update user", func(t *testing.T) {
		userData.On("SelectUserById", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		userData.On("EditUser", mock.AnythingOfType("user.UserCore")).Return(userValue, errors.New("error edit user")).Once()
		_, err := userUseCase.UpdateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error edit user")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("valid - delete user", func(t *testing.T) {
		userData.On("RemoveUser", mock.AnythingOfType("int")).Return(nil).Once()

		err := userUseCase.DeleteUser(userValue.ID)

		assert.Nil(t, err)
	})

	t.Run("invalid - error delete user", func(t *testing.T) {
		userData.On("RemoveUser", mock.AnythingOfType("int")).Return(errors.New("error remove user")).Once()

		err := userUseCase.DeleteUser(userValue.ID)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error remove user")
	})
}
