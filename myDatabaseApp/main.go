package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connnectToMySql() (*gorm.DB, error) {
	dsn := "pet_store:pet_store@tcp(localhost:3306)/pet_store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// In GORM, models represent database tables. Define your models as Go structs with appropriate field tags to map them to MySQL columns.
type PetStore struct {
	Id      uint 	`gorm:"primaryKey"`
	Address string 
	City    string 
	Name    string 	`gorm:"unique"`
	Phone   string 
	State   string 
	Zip     string 
}

// Learned From: 
// https://www.sqliz.com/posts/golang-gorm-mysql/
// https://gorm.io/docs/index.html
func main() {
	db, err := connnectToMySql()
	if err != nil {
		log.Fatal(err)
	}
	//defer db.close()
	log.Println("Step 1")
	// Perform database migration
	// This code ensures that the “User” table is created in the database.
	err = db.AutoMigrate(&PetStore{})
	if err != nil {
		log.Fatal(err)
	}

	// Your CRUD operations go here
	newPetStore := &PetStore{Address: "123 Main St.", City: "Atlanta", State: "GA", Phone: "3104445566", Zip: "11111", Name: "Wagging Tails"}
	err = createPetStore(db, newPetStore)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created PetStore:", newPetStore)

	petStoreID := newPetStore.Id
	petStore, err := getPetStoreByID(db, petStoreID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PetStore by ID:", petStore)

	//Update PetStore
	petStore.Phone = "8001111111"
	err = updatePetStore(db, petStore)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Update PetStore:", petStore)

	// Delete PetStore
	// err = deletePetStore(db, petStore)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Deleted PetStore:", petStore)
}

// CRUD operations
func createPetStore(db *gorm.DB, petStore *PetStore) error {
	result := db.Create(petStore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getPetStoreByID(db *gorm.DB, petStoreId uint) (*PetStore, error) {
	var petStore PetStore
	result := db.First(&petStore, petStoreId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &petStore, nil
}

func updatePetStore(db *gorm.DB, petStore *PetStore) error {
	result := db.Save(petStore)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deletePetStore(db *gorm.DB, petStore *PetStore) error {
	result := db.Delete(petStore)
	if result.Error != nil {
		return result.Error
	}
	return nil
}