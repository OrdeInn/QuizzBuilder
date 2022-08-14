package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizTemplate struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Questions []string           `json:"questions" bson:"questions"`
	UserId    string             `json:"userId" bson:"userId"`
}
type QuizInput struct {
	Questions []string `json:"questions" bson:"questions"`
	UserId    string   `json:"userId" bson:"userId"`
}
