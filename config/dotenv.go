package config

import (
	"log"

	"github.com/joho/godotenv"
)

type LoadENVService struct{}

func (e LoadENVService) NewLoadENV() LoadENVService {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file")
		// logger.LogError("SET-UP", "dplv2_command02", "Error loading .env file", "NewLoadENV()", err)
	}
	envVars := LoadENVService{}
	log.Printf("Env vars added")
	return envVars
}

func PrintEnvVariables() {
}
