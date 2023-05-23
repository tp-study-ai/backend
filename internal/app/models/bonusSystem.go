package models

import "time"

type Day struct {
	Day   time.Time `json:"day"`
	Count int       `json:"count_task"`
}

type Days struct {
	Days []Day `json:"days"`
}

type ShockMode struct {
	Today     bool `json:"today"`
	ShockMode int  `json:"chock_mode"`
}
