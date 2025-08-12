package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	initRoutes(router, db)
	return router
}

func initRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", healthCheckHandler)
	initCategoryRoutes(router, db)
	initStockRoutes(router, db)
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Go sauce démarrée 🚀"})
}
