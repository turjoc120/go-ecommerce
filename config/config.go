package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("PORT")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Error parsing PORT", err)
		os.Exit(1)
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() *Config {
	//makig sure config reads once
	if configurations == nil {
		LoadConfig()
	}
	return configurations
}
