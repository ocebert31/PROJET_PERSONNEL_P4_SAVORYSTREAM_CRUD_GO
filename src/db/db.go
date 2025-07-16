package db

import (
	"sauce-service/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"os"
	"github.com/joho/godotenv"
	"sauce-service/src/models"
)

type DBHandler interface {
	Close() error
	GetDB() *gorm.DB
}

type dbConnection struct {
	db *gorm.DB
}

func (d *dbConnection) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *dbConnection) GetDB() *gorm.DB {
	return d.db
}

// Types pour injection de dépendances
type GormOpener func(dsn string) (*gorm.DB, error)
type DSNProvider func() (string, error)

// ConnectWith est la fonction principale pour créer une connexion,
// avec injection des dépendances pour faciliter les tests et la flexibilité
func ConnectWith(getDSN DSNProvider, opener GormOpener) (DBHandler, error) {
	dsn, err := getDSN()
	if err != nil {
		return nil, err
	}
	database, err := opener(dsn)
	if err != nil {
		return nil, err
	}
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return &dbConnection{db: database}, nil
}

// Connect est la fonction utilisée par défaut,
// injecte config.GetDatabaseDSN et gorm.Open postgres
func Connect() (DBHandler, error) {
	return ConnectWith(config.GetDatabaseDSN, func(dsn string) (*gorm.DB, error) {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
}

func SetupTestDB(t *testing.T) *gorm.DB {
	err := godotenv.Load("../../.env.test")
    if err != nil {
        t.Fatalf("Erreur chargement fichier .env: %v", err)
    }
    dsn := os.Getenv("SAUCE_API_DSN_TEST")
    if dsn == "" {
        t.Fatalf("SAUCE_API_DSN_TEST non défini dans les variables d'environnement")
    }
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        t.Fatalf("failed to connect test db: %v", err)
    }
    db.Migrator().DropTable(&models.Category{})
    db.AutoMigrate(&models.Category{})
    return db
}