package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"sauce-service/src/app"
	"sauce-service/src/config"
	"sauce-service/src/db"
	"sauce-service/src/router"
	"sauce-service/src/server"
)

func main() {
	err := app.RunApp(configLoader{}, dbConnector{}, routerSetup{}, serverStarter{})
	if err != nil {
		log.Fatalf("Erreur au démarrage de l'app : %v", err)
	}
}

type configLoader struct{}
func (configLoader) LoadEnv() {
	config.LoadEnv()
}

type dbConnector struct{}
func (dbConnector) Connect() (app.DBConnection, error) {
	return db.Connect()
}

type routerSetup struct{}
func (routerSetup) Setup() *gin.Engine {
	return router.Setup()
}

type serverStarter struct{}
func (serverStarter) Start(engine *gin.Engine) error {
	return server.Server(engine) 
}
