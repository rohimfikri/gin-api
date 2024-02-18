package config

import "fmt"

func ConnectDBSys(ENV *Config) {
	fmt.Println("Connecting to Sys DB...")

	fmt.Println("CONNECTED to Sys DB!")
}

func DisconnectDBSys(ENV *Config) {
	fmt.Println("Disconnecting from Sys DB...")

	fmt.Println("DISCONNECTED from Sys DB!")
}
