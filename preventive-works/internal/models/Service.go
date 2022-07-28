package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	Name string             `json:"name" bson:"name"`
	Id   primitive.ObjectID `json:"id" bson:"_id"`
}
