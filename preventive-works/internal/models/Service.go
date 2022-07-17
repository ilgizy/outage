package models

type Service struct {
	Name string `json:"name" bson:"name"`
	Id   string `json:"id" bson:"id"`
}
