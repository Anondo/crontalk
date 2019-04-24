package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration file for the viper to parse
func LoadConfig() {
	if err := readFromBindedConfig(); err != nil {
		log.Println("Failed to read from binded config: ", err.Error())
		log.Println("Looking for local config file...")
		if err := readFromFileConfig(); err != nil {
			log.Fatal("Failed to read from local config: ", err.Error())
		}
	}
}

func readFromBindedConfig() error {
	cfgByte, err := Asset("confsdig.yml")
	if err != nil {
		return err
	}
	viper.SetConfigType("yml")
	if err := viper.ReadConfig(bytes.NewBuffer(cfgByte)); err != nil {
		return err
	}
	return nil
}

func readFromFileConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
