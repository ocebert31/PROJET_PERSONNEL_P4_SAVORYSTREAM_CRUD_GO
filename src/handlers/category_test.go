package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sauce-service/src/models"
	"sauce-service/src/app/category"
	"testing"
	"github.com/stretchr/testify/assert"
	test_utils "sauce-service/src/test_utils"
	"sauce-service/src/utils"
)

func TestCreateCategory(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	input := category.CategoryInput{
		Name: "Test Category",
	}
	jsonInput, _ := json.Marshal(input)

	resp := test_utils.PerformRequest(router, "POST", "/categories", jsonInput)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Category")
}

func TestCreateCategory_InvalidInput(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	body := []byte(`{"name":""}`)
	resp := test_utils.PerformRequest(router, "POST", "/categories", body)
	assert.NotEqual(t, http.StatusCreated, resp.Code)
}

func TestGetAllCategories(t *testing.T) {
	utils.SetupTestDB()
	utils.TestDB.Create(&models.Category{Name: "Cat1"})
	utils.TestDB.Create(&models.Category{Name: "Cat2"})
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	resp := test_utils.PerformRequest(router, "GET", "/categories", nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Cat1")
	assert.Contains(t, resp.Body.String(), "Cat2")
}

func TestGetCategoryByID_Success(t *testing.T) {
	utils.SetupTestDB()
	cat := models.Category{Name: "CatID"}
	utils.TestDB.Create(&cat)
	router := test_utils.SetupRouterForCategories(utils.TestDB)

	req, _ := http.NewRequest("GET", "/categories/"+cat.ID.String(), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "CatID")
}

func TestGetCategoryByID_NotFound(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	resp := test_utils.PerformRequest(router, "GET", "/categories/99999", nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestDeleteCategory_Success(t *testing.T) {
	utils.SetupTestDB()
	cat := models.Category{Name: "ToDelete"}
	utils.TestDB.Create(&cat)
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	resp := test_utils.PerformRequest(router, "DELETE", "/categories/"+cat.ID.String(), nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "deleted successfully")
}

func TestDeleteCategory_NotFound(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	resp := test_utils.PerformRequest(router, "DELETE", "/categories/99999", nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateCategory_Success(t *testing.T) {
	utils.SetupTestDB()
	cat := models.Category{Name: "OldName"}
	utils.TestDB.Create(&cat)
	router := test_utils.SetupRouterForCategories(utils.TestDB)

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
	utils.SetupTestDB()
	cat := models.Category{Name: "OldName"}
	utils.TestDB.Create(&cat)
	router := test_utils.SetupRouterForCategories(utils.TestDB)

	req, _ := http.NewRequest("PUT", "/categories/"+cat.ID.String(), bytes.NewBuffer([]byte(`{"name":""}`)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateCategory_NotFound(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForCategories(utils.TestDB)
	newInput := category.CategoryInput{Name: "UpdatedName"}
	jsonInput, _ := json.Marshal(newInput)
	resp := test_utils.PerformRequest(router, "PUT", "/categories/99999", jsonInput)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}