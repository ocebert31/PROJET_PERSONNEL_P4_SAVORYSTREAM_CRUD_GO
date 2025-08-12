package handlers_test

import (
	"encoding/json"
	"net/http"
	"sauce-service/src/models"
	"sauce-service/src/app/stock"
	"testing"
	"github.com/stretchr/testify/assert"
	test_utils "sauce-service/src/test_utils"
	utils "sauce-service/src/utils"
)

func TestCreateStock(t *testing.T) {
	utils.SetupTestDB()
	sauce, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce CreateStock")
	assert.NoError(t, err)
	router := test_utils.SetupRouterForStock(utils.TestDB)

	input := stock.CreateStockInput{
		SauceID:  sauce.ID.String(),
		Quantity: 10,
	}
	jsonInput, _ := json.Marshal(input)

	resp := test_utils.PerformRequest(router, "POST", "/stocks", jsonInput)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), `"quantity":10`)
	assert.Contains(t, resp.Body.String(), sauce.ID.String())
}

func TestCreateStock_InvalidInput(t *testing.T) {
	utils.SetupTestDB()
	router := test_utils.SetupRouterForStock(utils.TestDB)
	body := []byte(`{"quantity":5}`)
	resp := test_utils.PerformRequest(router, "POST", "/stocks", body)
	assert.NotEqual(t, http.StatusCreated, resp.Code)
}

func TestGetAllStocks(t *testing.T) {
	utils.SetupTestDB()
	sauce1, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce AllStocks 1")
	assert.NoError(t, err)
	sauce2, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce AllStocks 2")
	assert.NoError(t, err)
	utils.TestDB.Create(&models.Stock{SauceID: sauce1.ID, Quantity: 5})
	utils.TestDB.Create(&models.Stock{SauceID: sauce2.ID, Quantity: 12})

	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "GET", "/stocks", nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"quantity":5`)
	assert.Contains(t, resp.Body.String(), `"quantity":12`)
}

func TestGetStockByID_Success(t *testing.T) {
	utils.SetupTestDB()
	sauce, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce GetByID")
	assert.NoError(t, err)
	stock := models.Stock{SauceID: sauce.ID, Quantity: 20}
	utils.TestDB.Create(&stock)

	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "GET", "/stocks/"+stock.ID.String(), nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"quantity":20`)
}

func TestGetStockByID_NotFound(t *testing.T) {
	utils.SetupTestDB()
	fakeID := "3b6ddb86-8417-4896-aabd-236856fdab76"
	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "GET", "/stocks/"+fakeID, nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestDeleteStock_Success(t *testing.T) {
	utils.SetupTestDB()
	sauce, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce DeleteStock")
	assert.NoError(t, err)
	stock := models.Stock{SauceID: sauce.ID, Quantity: 15}
	utils.TestDB.Create(&stock)

	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "DELETE", "/stocks/"+stock.ID.String(), nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "deleted successfully")
}

func TestDeleteStock_NotFound(t *testing.T) {
	utils.SetupTestDB()
	fakeID := "bb22a94d-dd8e-458e-9f78-05e9683411a7"
	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "DELETE", "/stocks/"+fakeID, nil)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateStock_Success(t *testing.T) {
	utils.SetupTestDB()
	sauce, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce UpdateStock")
	assert.NoError(t, err)
	existingStock := models.Stock{SauceID: sauce.ID, Quantity: 8}
	utils.TestDB.Create(&existingStock)

	router := test_utils.SetupRouterForStock(utils.TestDB)
	newInput := stock.UpdateStockInput{Quantity: 42}
	jsonInput, _ := json.Marshal(newInput)

	resp := test_utils.PerformRequest(router, "PUT", "/stocks/"+existingStock.ID.String(), jsonInput)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), `"quantity":42`)
}

func TestUpdateStock_InvalidInput(t *testing.T) {
	utils.SetupTestDB()
	sauce, err := test_utils.CreateSauce(utils.TestDB, t, "Test Sauce UpdateStockInvalid")
	assert.NoError(t, err)
	stock := models.Stock{SauceID: sauce.ID, Quantity: 8}
	utils.TestDB.Create(&stock)

	router := test_utils.SetupRouterForStock(utils.TestDB)
	resp := test_utils.PerformRequest(router, "PUT", "/stocks/"+stock.ID.String(), []byte(`{"quantity":0}`))
	assert.NotEqual(t, http.StatusOK, resp.Code)
}

func TestUpdateStock_NotFound(t *testing.T) {
	utils.SetupTestDB()
	fakeID := "b6467030-2e46-4b0f-8604-49e12eb483ec"
	router := test_utils.SetupRouterForStock(utils.TestDB)

	newInput := stock.UpdateStockInput{Quantity: 33}
	jsonInput, _ := json.Marshal(newInput)

	resp := test_utils.PerformRequest(router, "PUT", "/stocks/"+fakeID, jsonInput)
	assert.NotEqual(t, http.StatusOK, resp.Code)
}