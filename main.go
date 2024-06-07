package main

import (
	"fmt"
	routes "github.com/venussa/Dealls-Technical-Test/app/routes"
	config "github.com/venussa/Dealls-Technical-Test/app/config"
	migration "github.com/venussa/Dealls-Technical-Test/database/migration"
	gin "github.com/gin-gonic/gin"
)

func main() {

	// load config
	getConfig := config.Config()

	// start auto migrate
	migration.Migration()

	// load routes
	router := gin.Default()
	routes.SetupAuthRoutes(router)
	router.Run(fmt.Sprintf(":%d", getConfig.Server.Port))
}