package question

import "go.mongodb.org/mongo-driver/bson/primitive"

// Question : The struct to map to question collection
type User struct {
	ID         string `bson:"_id,omitempty"`
	Answer     string `bson:"answer"`
	QuestionNo string `bson:"questionNo"`
	ImageUrl   string `bson:"imgUrl"`
}