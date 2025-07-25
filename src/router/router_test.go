package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// performRequest est une fonction utilitaire qui permet de simuler une requête HTTP
// à une route donnée sur un router Gin, et de capturer la réponse.
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// TestSetupRootRoute vérifie que la route GET "/" répond avec le bon message JSON et le code HTTP 200
func TestSetupRootRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := Setup(nil)
	w := performRequest(router, http.MethodGet, "/")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"API Go sauce démarrée 🚀"}`, w.Body.String())
}

func TestInitRoutesCallsInitCategoryRoutes(t *testing.T) {
    router := gin.New()
    db := &gorm.DB{}
    initRoutes(router, db)
    r := router.Routes()
    found := false
    for _, route := range r {
        if route.Path == "/categories" {
            found = true
            break
        }
    }
    if !found {
        t.Error("initCategoryRoutes semble ne pas avoir été appelée, route /categories absente")
    }
}
