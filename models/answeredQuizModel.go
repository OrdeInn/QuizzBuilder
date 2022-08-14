package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AnsweredQuiz struct {
	Id           primitive.ObjectID `json:"id" bson:"id"`
	ActionUserId string             `json:"actionUserId" bson:"actionUserId"`
	Answers      map[string]string  `json:"quizData" bson:"answers"`
}
