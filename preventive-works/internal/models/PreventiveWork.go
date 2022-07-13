package models

import "time"

type PreventiveWork struct {
	Id          int       `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	Deadline    time.Time `json:"deadline"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CountEvent  int       `json:"count_event"`
	IdService   int       `json:"id_service"`
	Events      []Event   `json:"events"`
}
