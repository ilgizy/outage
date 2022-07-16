package models

import "time"

type Event struct {
	Id               int       `json:"id" bson:"id"`
	CreateAt         time.Time `json:"create_at" bson:"create_at"`
	Deadline         time.Time `json:"deadline" bson:"deadline"`
	Description      string    `json:"description" bson:"description"`
	Status           string    `json:"status" bson:"status"`
	IdPreventiveWork int       `json:"id_preventive_work" bson:"id_preventive_work"`
}
