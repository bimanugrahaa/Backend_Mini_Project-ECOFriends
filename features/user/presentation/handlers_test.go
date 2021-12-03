package presentation_test

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"Backend_Mini_Project-ECOFriends/features/user/mocks"
	"Backend_Mini_Project-ECOFriends/features/user/presentation"
	"Backend_Mini_Project-ECOFriends/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
)

var (
	userUseCase mocks.Bussiness
	userHandler *presentation.UserHandler
	usersValue  []user.UserCore

	framework *echo.Echo
)

func TestMain(m *testing.M) {
	framework = routes.New()

	userHandler = presentation.NewUserHandler(&userUseCase)

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
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	echoContext := framework.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	t.Run("valid - get all user", func(t *testing.T) {

		userUseCase.On("GetAllUser").Return(usersValue).Once()

		assert.Equal(t, userHandler.GetAllUser(echoContext), nil)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
