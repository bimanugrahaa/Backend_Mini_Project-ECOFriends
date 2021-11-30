package routes

import (
	"Backend_Mini_Project-ECOFriends/config"
	"Backend_Mini_Project-ECOFriends/factory"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	auth := e.Group("")
	// auth.Use(middleware.JWT([]byte(config.JwtSecret)))

	//Donations

	e.GET("/donations/:id", presenter.DonationPresentation.GetDonationsById)
	e.POST("/donations", presenter.DonationPresentation.CreateDonation)
	e.DELETE("/donations/:id", presenter.DonationPresentation.DeleteDonationsById)
	e.PUT("/donations", presenter.DonationPresentation.UpdateDonation)

	e.POST("/login", presenter.UserPresentation.Login)

	e.POST("/donations/:id/comment", presenter.DonationPresentation.CreateComment)
	e.PUT("/donations/:id/comment", presenter.DonationPresentation.UpdateComment)
	e.DELETE("/donations/:id/comment", presenter.DonationPresentation.DeleteComment)
	//Users
	e.GET("/users", presenter.UserPresentation.GetAllUser)
	e.POST("/users", presenter.UserPresentation.CreateUser)
	e.PUT("/users", presenter.UserPresentation.UpdateUser)
	e.DELETE("/users/:id", presenter.UserPresentation.DeleteUser)

	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(config.JwtSecret),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"status":  "error login",
				"message": e,
			})
		},
	}))
	auth.GET("/donations", presenter.DonationPresentation.GetAllDonation)

	return e
}
