package models

import "time"

// The asterisk (*) in a struct field declaration indicates that the field is a pointer to the specified type.
// This usage is common when you want to represent an optional field in your struct, where the absence of a value can be indicated by a nil pointer.

// validate:"required": This part also seems to be a struct tag, but it's not a standard tag used by the encoding/json package. 
// Instead, it suggests that this struct may be used in conjunction with a validation library (like github.com/go-playground/validator).

// Remember, each variable needs to start with a Capitalized letter if you want it to be exported, and able to use in other packages
type Food struct {
	ID			uint 		`json:"food_id"` 
	Name       	*string  	`json:"name" validate:"required,min=2,max=100"`
	Price      	*float64 	`json:"price" validate:"required"`
	Food_image 	*string  	`json:"food_image"`
	Created_at 	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
	// `json:"-"` --- Indicates that this field should be excluded from the JSON output entirely
	// Food_id		string		`json:"food_id"`
	Menu_id		string		`json:"menu_id"`
}