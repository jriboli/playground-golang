package models

import "time"

type Table struct {
	ID					uint		`json:"table_id"`
	Number_of_guests	uint		`json:"number_of_guests" validate:"required"`
	Table_number		*int		`json:"table_number" validate:"required"`
	Created_at			time.Time 	`json:"created_at"`
	Updated_at			time.Time	`json:"updated_at"`
}