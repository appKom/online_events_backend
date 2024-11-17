package main

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

func LoadConfig() Config {
	return Config{
		DBHost:     os.Getenv("SUPABASE_DB_HOST"),
		DBPort:     os.Getenv("SUPABASE_DB_PORT"),
		DBUser:     os.Getenv("SUPABASE_DB_USER"),
		DBPassword: os.Getenv("SUPABASE_DB_PASSWORD"),
		DBName:     os.Getenv("SUPABASE_DB_NAME"),
		SSLMode:    "require",
	}
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.SSLMode)
}
