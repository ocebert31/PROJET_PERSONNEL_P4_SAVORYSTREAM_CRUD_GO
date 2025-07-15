package services

import (
	"sauce-service/src/models"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, name string) (*models.Category, error) {
	category := models.Category{Name: name}
	if err := db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllCategories(db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	if err := db.Order("created_at DESC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(db *gorm.DB, id string) (*models.Category, error) {
	var category models.Category
	if err := db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func DeleteCategory(db *gorm.DB, id string) error {
	category, err := GetCategoryByID(db, id)
	if err != nil {
		return err
	}
	return db.Delete(category).Error
}

func UpdateCategory(db *gorm.DB, id, name string) (*models.Category, error) {
	category, err := GetCategoryByID(db, id)
	if err != nil {
		return nil, err
	}
	category.Name = name
	if err := db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
