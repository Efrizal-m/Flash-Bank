package controllers

import (
	"flashbank/database"
	"flashbank/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSaldoByCustomerId(c *gin.Context) {
	var (
		result gin.H
	)
	customer_id, _ := strconv.Atoi(c.Param("customer_id"))

	customer_saldo, err := repository.GetCustomerLastSaldo(database.DbConnection, customer_id)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": customer_saldo,
		}
	}
	c.JSON(http.StatusOK, result)
}
