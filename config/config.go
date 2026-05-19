package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Host   string
	User   string
	Pass   string
	DbName string
	Port   string
}

type Config struct {
	Database Database
}

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Config{
		Database: Database{
			Host:   os.Getenv("DB_HOST"),
			Pass:   os.Getenv("DB_PASS"),
			User:   os.Getenv("DB_USER"),
			DbName: os.Getenv("DB_NAME"),
			Port:   os.Getenv("DB_PORT"),
		},
	}
}
