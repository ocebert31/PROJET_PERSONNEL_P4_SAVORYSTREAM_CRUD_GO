package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"errors"
)

const defaultPort = ":8080"

func LoadEnv() error {
    if err := godotenv.Load(); err != nil {
        log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement")
        return err
    }
    return nil
}

func GetAPIPort() string {
	if port := os.Getenv("SAUCE_API_PORT"); port != "" {
		return port
	}
	return defaultPort
}

func GetDatabaseDSN() (string, error) {
	dsn := os.Getenv("SAUCE_API_DSN")
	if dsn == "" {
		return "", errors.New("la variable SAUCE_API_DSN doit être définie")
	}
	return dsn, nil
}