package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envVariable(key string) string {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}

	return value
}

func main() {
	envValue := envVariable("API_KEY")

	fmt.Printf("env : %s = %s \n", "API_KEY", envValue)
}
