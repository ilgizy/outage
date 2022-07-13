package models

import "time"

type Event struct {
	Id               int       `json:"id"`
	CreateAt         time.Time `json:"create_at"`
	Deadline         time.Time `json:"deadline"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	IdPreventiveWork int       `json:"id_preventive_work"`
}
