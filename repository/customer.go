package repository

import (
	"database/sql"
	"flashbank/structs"
)

func GetAllCustomer(db *sql.DB) (results []structs.Customer, err error) {
	sql := "SELECT * FROM customer"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var customer = structs.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.IDCardNumber, &customer.Address, &customer.CIF, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, customer)
	}
	return
}

func GetCustomerById(db *sql.DB, id int) (customer structs.Customer, err error) {
	sql := "SELECT * FROM customer WHERE customer_id = $1"
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Name, &customer.IDCardNumber, &customer.Address, &customer.CIF, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			panic(err)
		}
	}

	return
}

func InsertCustomer(db *sql.DB, customer structs.Customer) (customer_id int, err error) {
	sql := "INSERT INTO customer (name, idcard_number, address, cif, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING customer_id"
	err = db.QueryRow(sql, customer.Name, customer.IDCardNumber, customer.Address, customer.CIF, customer.CreatedAt, customer.UpdatedAt).Scan(&customer.ID)
	customer_id = customer.ID
	return
}

func UpdateCustomer(db *sql.DB, customer structs.Customer) (err error) {
	sql := "UPDATE customer SET name = $2, idcard_number = $3, address = $4, cif = $5, updated_at = $6 WHERE customer_id = $1"
	errs := db.QueryRow(sql, customer.ID, customer.Name, customer.IDCardNumber, customer.Address, customer.CIF, customer.UpdatedAt)
	return errs.Err()
}

func DeleteCustomer(db *sql.DB, customer structs.Customer) (err error) {
	sql := "DELETE FROM customer WHERE customer_id = $1"
	errs := db.QueryRow(sql, customer.ID)
	return errs.Err()
}
