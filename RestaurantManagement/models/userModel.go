package models

import "time"

type User struct {
	ID				uint		`json:"user_id"`
	First_name		string		`json:"first_name"`
	Last_name		string		`json:"last_name"`
	Password		string		`json:"password"`
	Email			string		`json:"email"`
	Avatar			string		`json:"avatar"`
	Phone			string		`json:"phone"`
	Token			string		`json:"token"`
	Refresh_Token	string		`json:"refresh_token"`
	Created_at		time.Time	`json:"created_at"`
	Updated_at		time.Time	`json:"updated_at"`
}