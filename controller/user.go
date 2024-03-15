package controllers

import (
	"flashbank/database"
	"flashbank/repository"
	"flashbank/structs"
	"flashbank/utils"
	"net/http"
	"time"

	// log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userAdmin structs.UserAdmin

	if err := c.ShouldBindJSON(&userAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userAdmin.CreatedAt = time.Now()
	userAdmin.UpdatedAt = time.Now()
	err := repository.Register(database.DbConnection, userAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {
	var userAdmin structs.UserAdmin

	if err := c.ShouldBindJSON(&userAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAdmin_id, err := repository.Login(database.DbConnection, userAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or password is incorrect",
		})
		return
	}

	token, err := utils.GenerateToken(uint(userAdmin_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
