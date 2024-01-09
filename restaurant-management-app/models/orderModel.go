package models

import "time"

type Order struct {
	ID				uint		`json:"order_id"`
	Order_date		time.Time	`json:"order_date" validate:"required"`
	Created_at		time.Time	`json:"created_at"`
	Updated_at		time.Time	`json:"updated_at"`
	Table_id		string		`json:"table_id"`
}