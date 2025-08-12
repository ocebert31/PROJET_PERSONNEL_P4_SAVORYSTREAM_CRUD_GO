package router

import (
	"sauce-service/src/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initStockRoutes(router *gin.Engine, db *gorm.DB) {
	stock := router.Group("/stocks")
	{
		stock.POST("", handlers.CreateStock(db))
		stock.GET("", handlers.GetAllStocks(db))
		stock.GET("/:id", handlers.GetStockByID(db))
		stock.DELETE("/:id", handlers.DeleteStock(db))
		stock.PUT("/:id", handlers.UpdateStock(db))
	}
}
