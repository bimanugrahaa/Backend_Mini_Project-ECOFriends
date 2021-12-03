package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `"host":"${host}","method":"${method}","uri":"${uri}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println(string(reqBody))
		log.Println(string(resBody))
	}))

}
