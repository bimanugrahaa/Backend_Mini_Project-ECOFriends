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

	//Donations
	e.GET("/donations", presenter.DonationPresentation.GetAllDonation)
	e.GET("/donations/:id", presenter.DonationPresentation.GetDonationsById)
	e.GET("/donations/trending", presenter.DonationPresentation.GetDonationTrending)
	e.GET("/donations/latest", presenter.DonationPresentation.GetDonationLatest)
	auth.POST("/donations", presenter.DonationPresentation.CreateDonation)
	auth.DELETE("/donations/:id", presenter.DonationPresentation.DeleteDonationsById)
	auth.PUT("/donations/:id/edit", presenter.DonationPresentation.UpdateDonation)
	auth.PUT("/donations/:id", presenter.DonationPresentation.UpdateDonationValue)

	//Comments
	auth.POST("/donations/:id/comment", presenter.DonationPresentation.CreateComment)
	auth.PUT("/donations/:id/comment", presenter.DonationPresentation.UpdateComment)
	auth.DELETE("/donations/:id/comment", presenter.DonationPresentation.DeleteComment)

	//Users
	e.POST("/login", presenter.UserPresentation.Login)
	e.GET("/users", presenter.UserPresentation.GetAllUser)
	e.POST("/users", presenter.UserPresentation.CreateUser)
	auth.PUT("/users", presenter.UserPresentation.UpdateUser)
	auth.DELETE("/users/:id", presenter.UserPresentation.DeleteUser)

	return e
}
