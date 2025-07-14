package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	initRoutes(router)
	return router
}

func initRoutes(router *gin.Engine) {
	router.GET("/", healthCheckHandler)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Go sauce démarrée 🚀"})
}
