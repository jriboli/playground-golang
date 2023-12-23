package models

import "time"

type Menu struct {
	ID			uint 		`json:"menu_id"`
	Name		*string		`json:"name" validate:"required"`
	Category	string		`json:"category" validate:"required"`
	Start_date	time.Time	`json:"start_date"`
	End_date	time.Time	`json:"end_date"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
}