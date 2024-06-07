package driver

import (
  "log"
  "fmt"
  mysql "gorm.io/driver/mysql"
  gorm "gorm.io/gorm"
  CoreConfig "github.com/venussa/Dealls-Technical-Test/config"
)

type Config struct {
    Database struct {
        Host	string
        Port	int
        User	string
        Pass	string
        Dbname	string
		Charset	string
    }
}

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// load database config
	dbConfig := CoreConfig.Config().Database

	// attempt mysql connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return DB
}