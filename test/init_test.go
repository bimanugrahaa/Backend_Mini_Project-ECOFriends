package test

import (
	"Backend_Mini_Project-ECOFriends/migrate"
	"Backend_Mini_Project-ECOFriends/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	echoHandler *echo.Echo

	dbConn *gorm.DB
)

func init() {
	echoHandler = routes.New()

	// dbConn = config.InitDB("ecofriends-test")
	migrate.AutoMigrate()
}
