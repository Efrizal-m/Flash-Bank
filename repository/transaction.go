package repository

import (
	"database/sql"
	"flashbank/structs"
)

func GetAllTransaction(db *sql.DB) (results []structs.Transaction, err error) {
	sql := "SELECT * FROM transaction"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transaction = structs.Transaction{}
		err = rows.Scan(&transaction.ID, &transaction.SaldoID, &transaction.Volume, &transaction.TransactionType, &transaction.TransactionDate)
		if err != nil {
			panic(err)
		}
		results = append(results, transaction)
	}
	return
}

func AddTransaction(db *sql.DB, transaction structs.Transaction) (transaction_id int, err error) {
	sql := "INSERT INTO transaction (saldo_id, volume, transaction_type, transaction_date) VALUES ($1, $2, $3, $4) RETURNING transaction_id"
	err = db.QueryRow(sql, transaction.SaldoID, transaction.Volume, transaction.TransactionType, transaction.TransactionDate).Scan(&transaction.ID)
	transaction_id = transaction.ID
	return
}
