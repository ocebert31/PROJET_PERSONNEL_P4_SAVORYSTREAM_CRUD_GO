package app

import (
	"log"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type ConfigLoader interface {
	LoadEnv()
}

type DBConnection interface {
	Close() error
	GetDB() *gorm.DB
}

type DBConnector interface {
	Connect() (DBConnection, error)
}

type Router interface {
	Setup() *gin.Engine
}

type Server interface {
	Start(*gin.Engine) error
}

func RunApp( configLoader ConfigLoader, dbConnector DBConnector, router Router, server Server ) error {
	configLoader.LoadEnv()
	database, err := dbConnector.Connect()
	if err != nil {
		return err
	}
	defer func() {
		if cerr := database.Close(); cerr != nil {
			log.Printf("Erreur fermeture DB : %v", cerr)
		}
	}()
	db := database.GetDB()
	_ = db
	log.Println("✅ Connexion à PostgreSQL réussie")
	engine := router.Setup()
	return server.Start(engine)
}
