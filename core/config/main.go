package core_config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	fmt.Println("Loading app configuration...")

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }
	SetupViper()

	fmt.Println("App configuration has been LOADED!")
}

func SetupViper() {
	fmt.Println("Setup Viper...")

	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("Viper::Config file not found; ignore error if desired")
		} else {
			// Config file was found but another error was produced
			panic("Viper::" + err.Error())
		}
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		// Error load file config
		panic("Viper::" + err.Error())
	}

	// Config file found and successfully parsed
	fmt.Println("Viper setup has been DONE!")
}
