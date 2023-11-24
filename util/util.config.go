package util

import (
	"github.com/spf13/viper"
	"log"
)

var ApplicationConfig *viper.Viper

// LoadConfig loads the configuration based on the provided profile
func LoadConfig(activeProfilePath string) {
	v := viper.New()
	v.SetConfigName(activeProfilePath)
	v.AddConfigPath(".") // current directory

	// Read the configuration file corresponding to the activeProfilePath
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading the configuration file: %v", err)
	}

	ApplicationConfig = v
}
