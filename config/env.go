package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	PORT string
	DB   string
)

func LoadEnv() {
	loadEnvFile()

	PORT = getEnv("PORT", "5000")
	DB = getEnv("DB", "trello.db")
}

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}

func loadEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyValue := strings.Split(scanner.Text(), "=")
		os.Setenv(keyValue[0], keyValue[1])
	}
}
