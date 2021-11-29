package presentation

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	user_request "Backend_Mini_Project-ECOFriends/features/user/presentation/request"
	user_response "Backend_Mini_Project-ECOFriends/features/user/presentation/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBussiness user.Bussiness
}

func NewUserHandler(ubu user.Bussiness) *UserHandler {
	return &UserHandler{
		userBussiness: ubu,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {

	newUser := user_request.User{}

	c.Bind(&newUser)

	result, err := uh.userBussiness.CreateUser(user_request.ToCore(newUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCore(result),
	})
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	UpdateUser := user_request.User{}

	c.Bind(&UpdateUser)

	result, err := uh.userBussiness.UpdateUser(user_request.ToCore(UpdateUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCore(result),
	})
}

func (uh *UserHandler) GetAllUser(c echo.Context) error {
	result := uh.userBussiness.GetAllUser()
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "claims":  middleware.ExtractClaim(c),
		"message": "Success",
		"data":    user_response.FromCoreSlice(result),
	})
}

// func (uh *UserHandler)  {

// }
