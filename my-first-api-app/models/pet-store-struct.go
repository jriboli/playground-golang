package models

type PetStore struct {
	Id      string `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}