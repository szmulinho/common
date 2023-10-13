package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

type StorageConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
}

func (c StorageConfig) ConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.User, c.Password, c.Dbname, c.Port)
	return connectionString
}
