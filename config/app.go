package config

import "github.com/spf13/viper"

// AppConfig represents the application configuration
type AppConfig struct {
	Port    int
	Version string
}

var appCfg AppConfig

// LoadApp populates the AppConfig instance
func LoadApp() {
	appCfg = AppConfig{
		Port:    viper.GetInt("app.port"),
		Version: viper.GetString("app.version"),
	}
}

// App returns the AppConfig instance
func App() AppConfig {
	return appCfg
}
