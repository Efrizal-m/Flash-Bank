package controllers

import (
	"flashbank/database"
	"flashbank/repository"
	"flashbank/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	var transaction structs.Transaction
	var saldo structs.Saldo
	var report structs.Report

	customer_id, _ := strconv.Atoi(c.Param("customer_id"))

	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		panic(err)
	}

	saldo, err = repository.GetCustomerLastSaldo(database.DbConnection, customer_id)
	if err != nil {
		panic(err)
	}

	if transaction.TransactionType == "deposit" {
		saldo.Saldo = saldo.Saldo + transaction.Volume
		report.VolumeIn = transaction.Volume
	} else if transaction.TransactionType == "withdraw" {
		if saldo.Saldo > transaction.Volume {
			saldo.Saldo = saldo.Saldo - transaction.Volume
			report.VolumeOut = transaction.Volume
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Can't withdraw exceeds remaining saldo",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid transaction type input",
		})
		return
	}
	saldo.CustomerID = customer_id
	saldo.TransactionDate = time.Now()

	saldo_id, err := repository.AddSaldo(database.DbConnection, saldo)
	if err != nil {
		panic(err)
	}
	if saldo_id == 0 {
		saldo_id++
	}

	transaction.TransactionDate = time.Now()
	transaction.SaldoID = saldo_id

	transaction_id, err := repository.AddTransaction(database.DbConnection, transaction)
	if err != nil {
		panic(err)
	}
	if transaction_id == 0 {
		transaction_id++
	}

	report.TransactionID = transaction_id
	report.TransactionDate = time.Now()
	err = repository.AddReport(database.DbConnection, report)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert transaction",
	})
}
