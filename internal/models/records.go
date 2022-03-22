package models

import "time"

type Records struct {
	Counts    []int     `json:"counts"`
	CreatedAt time.Time `json:"createdAt"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}
