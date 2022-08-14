package routes

import (
	"github.com/gin-gonic/gin"
	"orderinn/QuizzBuilder/controllers"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/quizzes/:userId", controllers.GetUserQuizzes())
	router.POST("/quizzes", controllers.AddNewQuiz())
}
