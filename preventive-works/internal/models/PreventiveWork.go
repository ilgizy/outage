package models

import "time"

type PreventiveWork struct {
	Id          int       `json:"id" bson:"id"`
	CreateAt    time.Time `json:"create_at" bson:"create_at"`
	Deadline    time.Time `json:"deadline" bson:"deadline"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	CountEvent  int       `json:"count_event" bson:"count_event"`
	IdService   int       `json:"id_service" bson:"id_service"`
	Events      []Event   `json:"events" bson:"events"`
}
