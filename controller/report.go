package controllers

import (
	"flashbank/database"
	"flashbank/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReportByDate(c *gin.Context) {
	var (
		result gin.H
	)
	dateString := c.Param("tx_date")
	layout := "2006-01-02" // Layout for YYYY-MM-DD

	_, err := time.Parse(layout, dateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid transaction date input params",
		})
		return
	}
	customer, err := repository.GetReportByDate(database.DbConnection, dateString)

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
