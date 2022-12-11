package main

import (
	"golang-jwt/controllers"
	"golang-jwt/database"
	"golang-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	loadAppConfig()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1/")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
