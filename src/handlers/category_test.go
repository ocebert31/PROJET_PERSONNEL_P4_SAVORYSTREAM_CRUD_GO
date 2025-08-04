package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sauce-service/src/handlers"
	"sauce-service/src/models"
	"sauce-service/src/app/category"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"sauce-service/src/db"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/categories", handlers.CreateCategory(db))
	r.GET("/categories", handlers.GetAllCategories(db))
	r.GET("/categories/:id", handlers.GetCategoryByID(db))
	r.DELETE("/categories/:id", handlers.DeleteCategory(db))
	r.PUT("/categories/:id", handlers.UpdateCategory(db))
	return r
}

func performRequest(router http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestCreateCategory(t *testing.T) {
	db := db.SetupTestDB(t)
	router := setupRouter(db)
	input := category.CategoryInput{
		Name: "Test Category",
	}
	jsonInput, _ := json.Marshal(input)


	resp := performRequest(router, "POST", "/categories", jsonInput)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Category")
}

func TestCreateCategory_InvalidInput(t *testing.T) {
	db := db.SetupTestDB(t)
	router := setupRouter(db)
	body := []byte(`{"name":""}`)
	resp := performRequest(router, "POST", "/categories", body)
	assert.NotEqual(t, http.StatusCreated, resp.Code)
}

func TestGetAllCategories(t *testing.T) {
	db := db.SetupTestDB(t)
	db.Create(&models.Category{Name: "Cat1"})
	db.Create(&models.Category{Name: "Cat2"})
	router := setupRouter(db)
	resp := performRequest(router, "GET", "/categories", nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Cat1")
	assert.Contains(t, resp.Body.String(), "Cat2")
}

func TestGetCategoryByID_Success(t *testing.T) {
	db := db.SetupTestDB(t)
	cat := models.Category{Name: "CatID"}
	db.Create(&cat)
	router := setupRouter(db)

	req, _ := http.NewRequest("GET", "/categories/"+cat.ID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "CatID")
}

func TestGetCategoryByID_NotFound(t *testing.T) {
	db := db.SetupTestDB(t)
	router := setupRouter(db)
	resp := performRequest(router, "GET", "/categories/99999", nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestDeleteCategory_Success(t *testing.T) {
	db := db.SetupTestDB(t)
	cat := models.Category{Name: "ToDelete"}
	db.Create(&cat)
	router := setupRouter(db)
	resp := performRequest(router, "DELETE", "/categories/"+cat.ID.String(), nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "deleted successfully")
}

func TestDeleteCategory_NotFound(t *testing.T) {
	db := db.SetupTestDB(t)
	router := setupRouter(db)
	resp := performRequest(router, "DELETE", "/categories/99999", nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateCategory_Success(t *testing.T) {
	db := db.SetupTestDB(t)
	cat := models.Category{Name: "OldName"}
	db.Create(&cat)
	router := setupRouter(db)

	newInput := category.CategoryInput{Name: "UpdatedName"}
	jsonInput, _ := json.Marshal(newInput)

	req, _ := http.NewRequest("PUT", "/categories/"+cat.ID.String(), bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "UpdatedName")
}

func TestUpdateCategory_InvalidInput(t *testing.T) {
	db := db.SetupTestDB(t)
	cat := models.Category{Name: "OldName"}
	db.Create(&cat)
	router := setupRouter(db)

	req, _ := http.NewRequest("PUT", "/categories/"+cat.ID.String(), bytes.NewBuffer([]byte(`{"name":""}`)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateCategory_NotFound(t *testing.T) {
	db := db.SetupTestDB(t)
	router := setupRouter(db)
	newInput := category.CategoryInput{Name: "UpdatedName"}
	jsonInput, _ := json.Marshal(newInput)
	resp := performRequest(router, "PUT", "/categories/99999", jsonInput)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}
