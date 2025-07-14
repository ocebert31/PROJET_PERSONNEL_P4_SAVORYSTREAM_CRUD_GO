package server

import (
	"log"
	"sauce-service/src/config"
)

type Router interface {
	Run(...string) error
}

func Server(router Router) error {
	port := GetServerPort()
	log.Printf("Serveur démarré sur le port %s", port)
	return router.Run(port)
}

func GetServerPort() string {
	return config.GetAPIPort()
}