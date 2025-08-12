package testingUtils

import (
	"sauce-service/src/models"
	"testing"
	"gorm.io/gorm"
)

func CreateSauce(db *gorm.DB, t *testing.T, name string) (models.Sauce, error) {
	var category models.Category
	result := db.FirstOrCreate(&category, models.Category{Name: "Test Category"})
	if result.Error != nil {
		t.Fatalf("Failed to create category: %v", result.Error)
	}
	var sauce models.Sauce
	result = db.FirstOrCreate(&sauce, models.Sauce{
		Name:       name,
		CategoryID: category.ID,
	})
	if result.Error != nil {
		t.Fatalf("Failed to create sauce: %v", result.Error)
	}
	return sauce, nil
}