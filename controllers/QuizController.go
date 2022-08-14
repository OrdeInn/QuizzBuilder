package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"orderinn/QuizzBuilder/configs"
	"orderinn/QuizzBuilder/models"
	"time"
)

var quizCollection = configs.GetCollection(configs.DB, "quiz")

func GetUserQuizzes() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		opts := options.Find().SetSort(bson.D{{"id", 1}})
		cursor, err := quizCollection.Find(ctx, bson.D{{"userId", userId}}, opts)

		if err != nil {
			c.JSON(http.StatusNotFound,
				map[string]string{"message": err.Error()})
			return
		}

		var results []bson.M
		if err = cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError,
				map[string]string{"message": err.Error()})
			return
		}

		quizzes := make([]models.QuizTemplate, 0)
		for _, result := range results {
			var quiz models.QuizTemplate

			bsonBytes, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBytes, &quiz)

			if err = cursor.All(ctx, &results); err != nil {
				c.JSON(http.StatusInternalServerError,
					map[string]string{"message": err.Error()})
				return
			}

			quizzes = append(quizzes, quiz)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"error":    false,
			"quizList": quizzes,
		})
	}
}

func AddNewQuiz() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newQuiz models.QuizInput

		err := c.ShouldBindJSON(&newQuiz)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				map[string]string{"message": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := quizCollection.InsertOne(ctx, newQuiz)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				map[string]string{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"error": false,
			"data":  result,
		})
	}
}
