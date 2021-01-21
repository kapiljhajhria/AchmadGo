package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//Config func to get env value based on it's key/name prop
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

//GetAPIBase ...
func GetAPIBase() string {
	return Config("API_KEYWORD")
}

//Hash generates a hashed version of the password string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword checks if the supplied password is valid
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//Substring ...
func Substring(s string, start int, end int) string {
    startStrIdx := 0
    i := 0
    for j := range s {
        if i == start {
            startStrIdx = j
        }
        if i == end {
            return s[startStrIdx:j]
        }
        i++
    }
    return s[startStrIdx:]
}