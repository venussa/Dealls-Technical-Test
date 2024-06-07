package env

import (
	"fmt"
	"log"
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

	Server struct {
		Port	int
	}
}

func env() {
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
	serverConfig := cfg
	return serverConfig
}