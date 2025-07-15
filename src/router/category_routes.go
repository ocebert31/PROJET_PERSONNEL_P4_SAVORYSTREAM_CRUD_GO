package router

import (
	"sauce-service/src/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	category := router.Group("/categories")
	{
		category.POST("", handlers.CreateCategory(db))
		category.GET("", handlers.GetAllCategories(db))
		category.GET("/:id", handlers.GetCategoryByID(db))
		category.DELETE("/:id", handlers.DeleteCategory(db))
		category.PUT("/:id", handlers.UpdateCategory(db))
	}
}
