package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct { // We define a struct type name Config containing the allowed configuration
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *Config // var storing pointer of type Config

func LoadAppConfig() {
	log.Println("Loading Server Configuration...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
}
