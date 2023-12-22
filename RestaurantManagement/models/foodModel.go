package models

import "time"

type Food struct {
	ID			*float64 
	Name       	*string  	`json:"name" validate:"required,min=2,max=100"`
	Price      	*float64 	`json:"price" validate:"required"`
	Food_image 	*string  	`json:"food_image"`
	Created_at 	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
	Food_id		string		`json:"food_id"`
	Menu_id		string		`json:"menu_id"`
}