package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Variables struct {
	Host string
	Port string
	DBuser string
	Password string
	DBname string
	EchoUser string
	EchoPassword string
	EchoPostCode string
}

func ProjectEnvs() *Variables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return &Variables{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DBuser: os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
		DBname: os.Getenv("DB_NAME"),
		EchoUser: os.Getenv("ECHO_USER"),
		EchoPassword: os.Getenv("ECHO_PASSWORD"),
		EchoPostCode: os.Getenv("ECHO_POST_CODE"),
	}
}