package presentation

import (
	"Backend_Mini_Project-ECOFriends/features/user"
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
