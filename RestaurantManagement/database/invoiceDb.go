package database

import (
	"log"

	"github.com/binaryNomad/golang-learning/RestaurantManagement/models"
)

// CRUD operations
func CreateInvoice(invoice models.Invoice) (int64, error) {
    result, err := DB.Exec("INSERT INTO invoices (order_id, payment_method, payment_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", 
		invoice.Order_id, invoice.Payment_method, invoice.Payment_status, invoice.Created_at, invoice.Updated_at)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func GetInvoice(id int) (models.Invoice, error) {
    var invoice models.Invoice
    err := DB.QueryRow("SELECT id, payment_method, payment_status FROM invoices WHERE id = ?", id).
        Scan(&invoice.ID, &invoice.Payment_method, &invoice.Payment_status)

    if err != nil {
        return models.Invoice{}, err
    }

    return invoice, nil
}

func GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	rows, err := DB.Query("SELECT id, payment_method, payment_status FROM invoices")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var invoice models.Invoice
		err := rows.Scan(&invoice.ID, &invoice.Payment_method, &invoice.Payment_status)
		if err != nil {
			log.Fatal(err)
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil

}

func DeleteInvoice(id int) error {
    _, err := DB.Exec("DELETE FROM invoices WHERE id = ?", id)
    return err
}