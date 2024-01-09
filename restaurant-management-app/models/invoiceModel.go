package models

import "time"

type Invoice struct {
	ID 					uint 		`json:"invoice_id"`
	Order_id			uint		`json:"order_id"`
	Payment_method		string		`json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status		string		`json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date	time.Time	`json:"payment_due_date"`	
	Created_at			time.Time	`json:"created_at"`
	Updated_at			time.Time	`json:"updated_at"`
}