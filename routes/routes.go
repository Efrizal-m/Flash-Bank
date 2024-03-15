package routes

import (
	"net/http"

	controllers "flashbank/controller"
	middlewares "flashbank/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(router *gin.Engine) {
	router.GET("/", homeHandler)

	public := router.Group("/")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/adm")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/customer", controllers.GetAllCustomer)
	protected.POST("/customer", controllers.InsertCustomer)
	protected.GET("/customer/:id", controllers.GetCustomerById)
	protected.PUT("/customer/:id", controllers.Updatecustomer)
	protected.DELETE("/customer/:id", controllers.Deletecustomer)
	protected.POST("/transaction/:customer_id", controllers.AddTransaction)
	protected.GET("/saldo/:customer_id", controllers.GetSaldoByCustomerId)
	protected.GET("/report/:tx_date", controllers.GetReportByDate)
}

func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Flash Bank API")
}
