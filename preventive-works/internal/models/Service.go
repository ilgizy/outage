package models

type Service struct {
	Name string `json:"name" bson:"name"`
	Id   int    `json:"id" bson:"id"`
}
