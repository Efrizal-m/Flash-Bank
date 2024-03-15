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

	customer, err := repository.GetCustomerById(database.DbConnection, customer_id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": customer,
		}
	}
	c.JSON(http.StatusOK, result)
}
