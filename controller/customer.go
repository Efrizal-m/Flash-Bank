package controllers

import (
	"flashbank/database"
	"flashbank/repository"
	"flashbank/structs"

	token "flashbank/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCustomer(c *gin.Context) {
	var (
		result gin.H
	)

	_, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	customers, err := repository.GetAllCustomer(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": customers,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetCustomerById(c *gin.Context) {
	var (
		result gin.H
	)
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := repository.GetCustomerById(database.DbConnection, id)

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

func InsertCustomer(c *gin.Context) {
	var customer structs.Customer

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}
	customer.CIF = fmt.Sprintf(`%d-%s%s`, customer.ID, customer.IDCardNumber[:8], strings.ToUpper(customer.Address[:3]))
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	err = repository.InsertCustomer(database.DbConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert customer",
	})
}

func Updatecustomer(c *gin.Context) {
	var customer structs.Customer
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&customer)
	if err != nil {
		panic(err)
	}

	customer.ID = id
	customer.UpdatedAt = time.Now()
	customer.CIF = fmt.Sprintf(`%d-%s%s`, customer.ID, customer.IDCardNumber[:8], strings.ToUpper(customer.Address[:3]))

	err = repository.UpdateCustomer(database.DbConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update customer",
	})
}

func Deletecustomer(c *gin.Context) {
	var customer structs.Customer
	id, _ := strconv.Atoi(c.Param("id"))
	customer.ID = id

	err := repository.DeleteCustomer(database.DbConnection, customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete customer",
	})
}
