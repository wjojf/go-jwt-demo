package main

import (
	"golang-jwt/database"
)

func main() {
	loadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

}
