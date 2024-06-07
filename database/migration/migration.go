package migration

import (
	"log"
	driver "github.com/venussa/Dealls-Technical-Test/database/driver"
	models "github.com/venussa/Dealls-Technical-Test/models"
)

func Migration() {
	// connect database
	mysql := driver.ConnectDatabase()

	// start auto migrate
	err := mysql.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}