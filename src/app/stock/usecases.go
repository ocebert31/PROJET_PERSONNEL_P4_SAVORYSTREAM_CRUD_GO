package stock

import (
	"sauce-service/src/models"
	"sauce-service/src/services"
	"gorm.io/gorm"
)

func CreateStockFromInput(db *gorm.DB, input CreateStockInput) (*models.Stock, error) {
	return services.CreateStock(db, input.SauceID, input.Quantity)
}

func UpdateStockFromInput(db *gorm.DB, id string, input UpdateStockInput) (*models.Stock, error) {
	return services.UpdateStock(db, id, input.Quantity)
}
