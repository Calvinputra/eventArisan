package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func InitEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Could not load .env file. Empty:", err)
		log.Fatalf("Empty loading .env file")
	}
}

func GetStringEnv(key string) string {
	return os.Getenv(key)
}

func GetIntEnv(key string) int {
	num, numErr := strconv.Atoi(os.Getenv(key))

	if numErr != nil {
		panic(fmt.Sprintf("Empty converting env var %s to int", key))
	}

	return num
}
