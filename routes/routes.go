package routes

import (
	"Backend_Mini_Project-ECOFriends/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/donations", presenter.DonationPresentation.GetAllDonation)
	e.GET("/users", presenter.UserPresentation.GetAllUser)

	return e
}
