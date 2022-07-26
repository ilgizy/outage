package models

import "time"

type Event struct {
	CreateAt    time.Time `json:"create_at" bson:"create_at"`
	Deadline    time.Time `json:"deadline" bson:"deadline"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
}
