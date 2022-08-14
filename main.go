package main

import (
	"github.com/gin-gonic/gin"
	"orderinn/QuizzBuilder/configs"
	"orderinn/QuizzBuilder/routes"
)

func main() {
	configs.ConnectDB()

	router := gin.Default()

	routes.InitRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
