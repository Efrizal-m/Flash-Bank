package main

import (
	"database/sql"
	"flashbank/database"
	"flashbank/helper"
	"flashbank/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load environtment")
	} else {
		fmt.Println("succes load environtment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database Conenction Failed")
		panic(err)
	} else {
		fmt.Println("Database Conenction Success")
	}
	database.DBMigrate(DB)
	defer DB.Close()

	//Route
	router := gin.Default()
	routes.RegisterRoutes(router)

	port := helper.EnvPortOr("3000")
	err := router.Run(port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
