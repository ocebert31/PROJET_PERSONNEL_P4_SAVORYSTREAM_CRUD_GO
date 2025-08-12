package services

import (
	"sauce-service/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"fmt"
)

func CreateStock(db *gorm.DB, sauceID string, quantity int) (*models.Stock, error) {
	parsedID, err := uuid.Parse(sauceID)
	if err != nil {
		return nil, err
	}
	var existing models.Stock
	if err := db.Where("sauce_id = ?", parsedID).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("un stock existe déjà pour cette sauce")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err 
	}
	stock := models.Stock{
		SauceID:  parsedID,
		Quantity: quantity,
	}
	if err := db.Create(&stock).Error; err != nil {
		return nil, err
	}
	if err := db.Preload("Sauce").First(&stock, "id = ?", stock.ID).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	var stocks []models.Stock
	if err := db.Preload("Sauce").Order("created_at DESC").Find(&stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}

func GetStockByID(db *gorm.DB, id string) (*models.Stock, error) {
	var stock models.Stock
	if err := db.Preload("Sauce").First(&stock, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func DeleteStock(db *gorm.DB, id string) error {
	stock, err := GetStockByID(db, id)
	if err != nil {
		return err
	}
	return db.Delete(stock).Error
}

func UpdateStock(db *gorm.DB, id string, quantity int) (*models.Stock, error) {
	stock, err := GetStockByID(db, id)
	if err != nil {
		return nil, err
	}
	stock.Quantity = quantity
	if err := db.Preload("Sauce").Save(stock).Error; err != nil {
		return nil, err
	}
	return stock, nil
}