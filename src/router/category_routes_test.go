package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"sauce-service/src/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitCategoryRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := db.SetupTestDB(t) 
	router := gin.New()
	initCategoryRoutes(router, db)
	routes := []struct {
		method string
		path   string
	}{
		{"POST", "/categories"},
		{"GET", "/categories"},
		{"GET", "/categories/1"},
		{"DELETE", "/categories/1"},
		{"PUT", "/categories/1"},
	}
	for _, route := range routes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.NotEqual(t, http.StatusNotFound, w.Code, "La route %s %s ne devrait pas renvoyer 404", route.method, route.path)
	}
}
