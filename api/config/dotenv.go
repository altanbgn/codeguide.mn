package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile(envFile string) {
  error := godotenv.Load(envFile)

  if error != nil {
    log.Fatalf("can't load .env file. error %v", error)
  }
}
