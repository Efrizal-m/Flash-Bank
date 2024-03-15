package routes

import (
	"net/http"

	controllers "flashbank/controller"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(router *gin.Engine) {
	router.GET("/", homeHandler)
	router.GET("/about", aboutHandler)
	// Add more routes as needed

	router.GET("/customer", controllers.GetAllCustomer)
	router.POST("/customer", controllers.InsertCustomer)
	router.GET("/customer/:id", controllers.GetCustomerById)
	router.PUT("/customer/:id", controllers.Updatecustomer)
	router.DELETE("/customer/:id", controllers.Deletecustomer)

	router.GET("/saldo/:customer_id", controllers.GetSaldoByCustomerId)

	router.POST("/transaction/:customer_id", controllers.AddTransaction)
	router.GET("/report/:tx_date", controllers.GetReportByDate)
}

func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Home Page")
}

func aboutHandler(c *gin.Context) {
	c.String(http.StatusOK, "About Page")
}
