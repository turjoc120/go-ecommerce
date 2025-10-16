package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DbConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMode  bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DbConfig
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION not set in .env file")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in .env file")
		os.Exit(1)
	}

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		fmt.Println("PORT not set in .env file")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT_SECRET_KEY not set in .env file")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Error parsing PORT", err)
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("DB_HOST not set in .env file")
		os.Exit(1)
	}

	dbprt := os.Getenv("DB_PORT")
	if dbprt == "" {
		fmt.Println("DB_PORT not set in .env file")
		os.Exit(1)
	}

	dbPort, err := strconv.ParseInt(dbprt, 10, 64)
	if err != nil {
		fmt.Println("Error parsing PORT", err)
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB_USER not set in .env file")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD not set in .env file")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_DBNAME")
	if dbName == "" {
		fmt.Println("DB_DBNAME not set in .env file")
		os.Exit(1)
	}

	dbSslMode := os.Getenv("DB_SSLMode")
	enableSSLMode, err := strconv.ParseBool(dbSslMode)
	if err != nil {
		fmt.Println("Error parsing DB_SSLMode", err)
		os.Exit(1)
	}

	dbConfig := DbConfig{
		HOST:     dbhost,
		PORT:     int(dbPort),
		USER:     dbUser,
		PASSWORD: dbPassword,
		DBNAME:   dbName,
		SSLMode:  enableSSLMode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           &dbConfig,
	}

}

func GetConfig() *Config {
	//makig sure config reads once
	if configurations == nil {
		LoadConfig()
	}
	return configurations
}
