package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"sauce-service/src/utils"
)

func TestInitStockRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	utils.SetupTestDB()
    initStockRoutes(router, utils.TestDB)
	routes := []struct {
		method string
		path   string
	}{
		{"POST", "/stocks"},
		{"GET", "/stocks"},
		{"GET", "/stocks/1"},
		{"DELETE", "/stocks/1"},
		{"PUT", "/stocks/1"},
	}
	for _, route := range routes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusNotFound, w.Code, "La route %s %s ne devrait pas renvoyer 404", route.method, route.path)
	}
}
