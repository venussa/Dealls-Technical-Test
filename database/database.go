package database

import (
  "log"
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "login-api/models"
  "github.com/spf13/viper"
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

func ConnectDatabase() {
	// load database config
	viper.SetConfigFile("env.yaml")
	errConfig := viper.ReadInConfig()
    if errConfig != nil {
        log.Fatalf("Error reading config file: %s", errConfig)
    }
    var cfg Config
    errConfig = viper.Unmarshal(&cfg)
    if errConfig != nil {
        log.Fatalf("Unable to decode config file: %s", errConfig)
    }
	dbConfig := cfg.Database

	// attempt mysql connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Dbname)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

  // Automatically migrate the User model schema
  err = DB.AutoMigrate(&models.User{})
  if err != nil {
    log.Fatal("Failed to migrate database:", err)
  }

  log.Println("Database connection established and migration completed")
}