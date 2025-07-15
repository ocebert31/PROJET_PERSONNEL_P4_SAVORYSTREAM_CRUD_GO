package category

import (
	"sauce-service/src/models"
	"sauce-service/src/services"
	"gorm.io/gorm"
)

func CreateFromInput(db *gorm.DB, input CategoryInput) (*models.Category, error) {
	return services.CreateCategory(db, input.Name)
}

func UpdateFromInput(db *gorm.DB, id string, input CategoryInput) (*models.Category, error) {
	return services.UpdateCategory(db, id, input.Name)
}
