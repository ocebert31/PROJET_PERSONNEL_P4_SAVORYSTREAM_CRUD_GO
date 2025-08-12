package services_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sauce-service/src/models"
	"gorm.io/gorm"
	"sauce-service/src/utils"
	"sauce-service/src/services"
	test_utils "sauce-service/src/test_utils"
)

func createStock(t *testing.T, db *gorm.DB, sauceID string, quantity int) *models.Stock {
	stock, err := services.CreateStock(db, sauceID, quantity) 
	if err != nil {
		t.Fatalf("failed to create stock: %v", err)
	}
	return stock
}

func createSauceWithStock(t *testing.T, db *gorm.DB, sauceName string, quantity int) (*models.Stock, error) {
	sauce, err := test_utils.CreateSauce(db, t, sauceName)
	if err != nil {
		return nil, err
	}
	return createStock(t, db, sauce.ID.String(), quantity), nil
}

func TestCreateStock(t *testing.T) {
    t.Run("CreateStock_Success", func(t *testing.T) {
		utils.SetupTestDB()
		stock, err := createSauceWithStock(t, utils.TestDB, "Test Sauce CreateStock", 5)
		assert.NoError(t, err)
        assert.Equal(t, 5, stock.Quantity)
    })

    t.Run("CreateStock_EmptyQuantity", func(t *testing.T) {
		utils.SetupTestDB()
		stock, err := createSauceWithStock(t, utils.TestDB, "Sauce for EmptyQuantity test", 0)
        assert.NoError(t, err)
        assert.Equal(t, 0, stock.Quantity)
    })
}

func TestGetAllStocks(t *testing.T) {
    t.Run("GetAllStocks_Success", func(t *testing.T) {
		utils.SetupTestDB()
		createSauceWithStock(t, utils.TestDB, "Test Sauce1 GetAllStocks_Success", 6)
		createSauceWithStock(t, utils.TestDB, "Test Sauce2 GetAllStocks_Success", 5)
		stocks, err := services.GetAllStocks(utils.TestDB)
		assert.NoError(t, err)
		assert.Len(t, stocks, 2)
		quantities := []int{stocks[0].Quantity, stocks[1].Quantity}
        assert.Contains(t, quantities, 6)
        assert.Contains(t, quantities, 5)
    })

    t.Run("GetAllStocks_EmptyDB", func(t *testing.T) {
		utils.SetupTestDB()
		stocks, err := services.GetAllStocks(utils.TestDB)
		assert.NoError(t, err)
		assert.Len(t, stocks, 0)
    })
}

func TestGetStockByID(t *testing.T) {
	t.Run("GetStockByID_Success", func(t *testing.T) {
		utils.SetupTestDB()
		stock, err := createSauceWithStock(t, utils.TestDB, "Test Sauce GetStockByID_Success", 2)
		assert.NoError(t, err)
		stockById, err := services.GetStockByID(utils.TestDB, stock.ID.String())
		assert.NotNil(t, stockById)
		assert.Equal(t, stockById.Quantity, 2)
	})

	t.Run("GetStockByID_NotFound", func(t *testing.T) {
		utils.SetupTestDB()
		_, err := services.GetStockByID(utils.TestDB, "id-invalide")
		assert.Error(t, err)
	})
}

func TestUpdateStock(t *testing.T) {
	t.Run("UpdateStock_Success", func(t *testing.T) {
		utils.SetupTestDB()
		stock, err := createSauceWithStock(t, utils.TestDB, "Test Sauce UpdateStock_Success", 3)
		assert.NoError(t, err)
		updated, err := services.UpdateStock(utils.TestDB, stock.ID.String(), 6)
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, updated.Quantity, 6)
	})

	t.Run("UpdateStock_InvalidID", func(t *testing.T) {
		utils.SetupTestDB()
		updated, err := services.UpdateStock(utils.TestDB, "inexistant-id", 5)
		assert.Error(t, err)
		assert.Nil(t, updated)
	})
}

func TestDeleteStock(t *testing.T) {
	t.Run("DeleteStock_Success", func(t *testing.T) {
		utils.SetupTestDB()
		stock, err := createSauceWithStock(t, utils.TestDB, "Test Sauce DeleteStock_Success", 4)
		assert.NoError(t, err)
		errDelete := services.DeleteStock(utils.TestDB, stock.ID.String())
		assert.NoError(t, errDelete)
		_, err = services.GetStockByID(utils.TestDB, stock.ID.String())
		assert.Error(t, err)
	})

	t.Run("DeleteStock_InvalidID", func(t *testing.T) {
		utils.SetupTestDB()
		err := services.DeleteStock(utils.TestDB, "inexistant-id")
		assert.Error(t, err)
	})
}