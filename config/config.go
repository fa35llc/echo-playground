package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
Caution!  Add .env to .gitignore.
*/
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	e := godotenv.Load() //Load env.file
	if e != nil {
		log.Println(e)
	}
	config := Configuration{}
	config.DB_USERNAME = os.Getenv("DB_USERNAME")
	config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_NAME = os.Getenv("DB_NAME")

	return config
}
