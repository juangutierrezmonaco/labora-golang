package models

import "time"

type Movie struct {
	Id           *int       `json:"id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Genre        *string    `json:"genre,omitempty"`
	ReleaseDate  *time.Time `json:"release_date,omitempty"`
	AcquiredDate *time.Time `json:"acquired_date,omitempty"`
	Stock        *int       `json:"stock,omitempty"`
	Price        *float64   `json:"price,omitempty"`
}
