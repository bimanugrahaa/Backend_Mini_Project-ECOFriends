package main

import (
	"Backend_Mini_Project-ECOFriends/config"
	"Backend_Mini_Project-ECOFriends/migrate"
	"Backend_Mini_Project-ECOFriends/routes"
)

func main() {
	config.InitDB()
	migrate.AutoMigrate()

	e := routes.New()

	e.Start(":8000")
}
