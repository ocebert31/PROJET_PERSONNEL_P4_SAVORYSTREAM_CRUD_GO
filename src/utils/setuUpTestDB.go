package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sauce-service/src/models"
)

var TestDB *gorm.DB

func SetupTestDB() {
    // Chercher .env.test en remontant les dossiers
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Impossible de récupérer le répertoire courant: %v", err)
    }
    var envPath string
    for {
        testPath := filepath.Join(currentDir, ".env.test")
        if _, err := os.Stat(testPath); err == nil {
            envPath = testPath
            break
        }

        parent := filepath.Dir(currentDir)
        if parent == currentDir {
            log.Fatal(".env.test introuvable dans l'arborescence")
        }
        currentDir = parent
    }
    // Charger le fichier trouvé
    if err := godotenv.Load(envPath); err != nil {
        log.Fatalf("Erreur chargement fichier .env: %v", err)
    }
    // Récupérer la DSN
    dsn := os.Getenv("SAUCE_API_DSN_TEST")
    if dsn == "" {
        log.Fatal("SAUCE_API_DSN_TEST non défini")
    }
    // Connexion DB
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Connexion test DB échouée: %v", err)
    }
    // Migration
    if err := db.AutoMigrate(&models.Category{}, &models.Sauce{}, &models.Stock{}); err != nil {
        log.Fatalf("Erreur migration test DB: %v", err)
    }
    if err := db.Exec("TRUNCATE TABLE stocks, sauces, categories RESTART IDENTITY CASCADE").Error; err != nil {
        log.Fatalf("Erreur nettoyage tables test DB: %v", err)
    }
    TestDB = db
}
