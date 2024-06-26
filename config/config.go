package config

import (
	"log"
	viper "github.com/spf13/viper"
)

type Configs struct {
	
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

func Config() (result Configs) {
	viper.SetConfigFile("config.yaml")
	errConfig := viper.ReadInConfig()
    if errConfig != nil {
        log.Fatalf("Error reading config file: %s", errConfig)
    }
    var cfg Configs
    errConfig = viper.Unmarshal(&cfg)
    if errConfig != nil {
        log.Fatalf("Unable to decode config file: %s", errConfig)
    }
	
	result = cfg
	return result
}