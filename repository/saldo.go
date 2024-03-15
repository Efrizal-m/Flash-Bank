package repository

import (
	"database/sql"
	"flashbank/structs"
)

func GetSaldoByCustomer(db *sql.DB, customer_id int) (saldo structs.Saldo, err error) {
	sql := "SELECT * FROM saldo WHERE customer_id = $1"
	rows, err := db.Query(sql, customer_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&saldo.ID, &saldo.CustomerID, &saldo.Saldo, &saldo.TransactionDate)
		if err != nil {
			panic(err)
		}
	}
	return
}

func GetSaldoByCustomerAndDate(db *sql.DB, customer_id int, date string) (saldo structs.Saldo, err error) {
	sql := "SELECT * FROM saldo WHERE customer_id = $1 and DATE(transaction_date) = $2"
	rows, err := db.Query(sql, customer_id, date)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&saldo.ID, &saldo.CustomerID, &saldo.Saldo, &saldo.TransactionDate)
		if err != nil {
			panic(err)
		}
	}

	return
}

func AddSaldo(db *sql.DB, saldo structs.Saldo) (err error) {
	sql := "INSERT INTO saldo (customer_id, saldo, transaction_date) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, saldo.CustomerID, saldo.Saldo, saldo.TransactionDate)
	return errs.Err()
}

func GetCustomerLastSaldo(db *sql.DB, customer_id int) (saldo structs.Saldo, err error) {
	sql := `
		SELECT *
		FROM saldo s
		WHERE customer_id = $1
		ORDER BY saldo_id DESC
		LIMIT 1;
	`
	rows, err := db.Query(sql, customer_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&saldo.ID, &saldo.CustomerID, &saldo.Saldo, &saldo.TransactionDate)
		if err != nil {
			panic(err)
		}
	}

	return
}
