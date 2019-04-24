package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration file for the viper to parse
func LoadConfig() {
	cfgByte, err := Asset("config.yml")
	if err != nil {
		log.Fatal("Failed to read config file: ", err.Error())
	}
	viper.SetConfigType("yml")
	if err := viper.ReadConfig(bytes.NewBuffer(cfgByte)); err != nil {
		log.Fatal("Failed to read config file: ", err.Error())
	}
}
