package testingUtils

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http/httptest"
	"bytes"
	"sauce-service/src/handlers"
)

func PerformRequest(router http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func SetupRouterForCategories(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/categories", handlers.CreateCategory(db))
	r.GET("/categories", handlers.GetAllCategories(db))
	r.GET("/categories/:id", handlers.GetCategoryByID(db))
	r.DELETE("/categories/:id", handlers.DeleteCategory(db))
	r.PUT("/categories/:id", handlers.UpdateCategory(db))
	return r
}

func SetupRouterForStock(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/stocks", handlers.CreateStock(db))
	r.GET("/stocks", handlers.GetAllStocks(db))
	r.GET("/stocks/:id", handlers.GetStockByID(db))
	r.DELETE("/stocks/:id", handlers.DeleteStock(db))
	r.PUT("/stocks/:id", handlers.UpdateStock(db))
	return r
}