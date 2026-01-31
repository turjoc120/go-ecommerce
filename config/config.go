package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DBNAME   string
	SSLMode  bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	DB          *DbConfig
}

var configuration *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("service name is required")
		os.Exit(1)
	}
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("http port is required")
		os.Exit(1)
	}
	port, err := strconv.Atoi(httpPortStr)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("jwt secret is required")
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

	//port
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

	//host
	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("DB_HOST not set in .env file")
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

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		JwtSecret:   jwtSecret,
		DB:          &dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}

	return configuration
}
