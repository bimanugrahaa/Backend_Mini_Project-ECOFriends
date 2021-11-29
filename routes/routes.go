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

	//Donations
	e.GET("/donations", presenter.DonationPresentation.GetAllDonation)
	e.GET("/donations/:id", presenter.DonationPresentation.GetDonationsById)
	e.POST("/donations", presenter.DonationPresentation.CreateDonation)
	e.DELETE("/donations/:id", presenter.DonationPresentation.DeleteDonationsById)
	e.PUT("/donations", presenter.DonationPresentation.UpdateDonation)
	// e.POST("/donations", presenter.DonationPresentation.CreateDescriptionDonation)

	e.GET("/users", presenter.UserPresentation.GetAllUser)
	e.POST("/users", presenter.UserPresentation.CreateUser)
	e.PUT("/users", presenter.UserPresentation.UpdateUser)

	return e
}
