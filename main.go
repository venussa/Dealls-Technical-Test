package main

import (
	"fmt"
	"log"
	"login-api/database"
	"login-api/app/routes"
	"env"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	config := env()

	// connect database
	database.ConnectDatabase()

	// load routes
	router := gin.Default()
	routes.SetupAuthRoutes(router)
	router.Run(fmt.Sprintf(":%d", config.Server.Port))
}