package database

import (
	"log"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/models"
)

// CRUD operations
func CreateFood(food models.Food) (int64, error) {
    result, err := DB.Exec("INSERT INTO foods (name, price) VALUES (?, ?)", food.Name, food.Price)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func GetFood(id int) (models.Food, error) {
    var food models.Food
    err := DB.QueryRow("SELECT id, name, price FROM foods WHERE id = ?", id).
        Scan(&food.ID, &food.Name, &food.Price)

    if err != nil {
        return models.Food{}, err
    }

    return food, nil
}

func GetFoods() ([]models.Food, error) {
	var foods []models.Food
	rows, err := DB.Query("SELECT id, name, price FROM foods")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var food models.Food
		err := rows.Scan(&food.ID, &food.Name, &food.Price)
		if err != nil {
			log.Fatal(err)
		}
		foods = append(foods, food)
	}

	return foods, nil

}

func DeleteFood(id int) error {
    _, err := DB.Exec("DELETE FROM foods WHERE id = ?", id)
    return err
}