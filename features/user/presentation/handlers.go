package presentation

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	user_request "Backend_Mini_Project-ECOFriends/features/user/presentation/request"
	user_response "Backend_Mini_Project-ECOFriends/features/user/presentation/response"
	"Backend_Mini_Project-ECOFriends/middleware"

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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "email available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCore(result),
	})
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	UpdateUser := user_request.User{}
	claim := middleware.ExtractTokenUserId(c)
	user_id := int(claim["user_id"].(float64))

	UpdateUser.ID = user_id
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

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	claim := middleware.ExtractTokenUserId(c)
	user_id := int(claim["user_id"].(float64))

	err := uh.userBussiness.DeleteUser(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete user by id success",
	})
}

func (uh *UserHandler) GetAllUser(c echo.Context) error {
	result := uh.userBussiness.GetAllUser()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    user_response.FromCoreSlice(result),
	})
}

func (uh *UserHandler) Login(c echo.Context) error {
	infoUser := user_request.User{}

	c.Bind(&infoUser)

	result, err := uh.userBussiness.Login(user_request.ToCore(infoUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCoreLogin(result),
	})
}
