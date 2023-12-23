package models

import "time" 

type OrderItem struct {
	ID   			uint 		`json:"order_item_id"`
	Quantity		int			`json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Unit_price		int			`json:"unit_price"`
	Created_at		time.Time	`json:"created_at"`
	Updated_at		time.Time	`json:"updated_at"`
	Food_id			string		`json:"food_id"`
	Order_id		string		`json:"order_id"`
}