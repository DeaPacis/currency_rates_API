package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	Version string
	Service string
	Author  string
}

func Load() Config {
	godotenv.Load()

	port := getEnv("PORT", "8000")
	if !regexp.MustCompile(`^\d+$`).MatchString(port) {
		log.Println("Port value is invalid. Default 8000 is used")
		port = "8000"
	}

	service := getEnv("SERVICE", "currency")
	version := getEnv("VERSION", "1.0.0")
	author := getEnv("AUTHOR", "a.pochebut")

	return Config{
		Port:    port,
		Version: version,
		Service: service,
		Author:  author,
	}
}

func getEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}
