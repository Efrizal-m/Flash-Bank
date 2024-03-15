package repository

import (
	"database/sql"
	"flashbank/structs"
)

func GetReportByDate(db *sql.DB, tx_date string) (results []structs.Report, err error) {
	sql := "SELECT * FROM report WHERE DATE(transaction_date) = $1"

	rows, err := db.Query(sql, tx_date)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var report = structs.Report{}
		err = rows.Scan(&report.ID, &report.TransactionID, &report.VolumeIn, &report.VolumeOut, &report.TransactionDate)
		if err != nil {
			panic(err)
		}
		results = append(results, report)
	}
	return
}

func AddReport(db *sql.DB, report structs.Report) (err error) {
	sql := "INSERT INTO report (transaction_id, volume_in, volume_out, transaction_date) VALUES ($1, $2, $3, $4)"
	errs := db.QueryRow(sql, report.TransactionID, report.VolumeIn, report.VolumeOut, report.TransactionDate)
	return errs.Err()
}
