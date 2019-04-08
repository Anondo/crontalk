package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration file for the viper to parse
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config file")
	}
}
