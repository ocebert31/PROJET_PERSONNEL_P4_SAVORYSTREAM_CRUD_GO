package services_test

import (
	"testing"
	"sauce-service/src/services"
	"sauce-service/src/db"
	"github.com/stretchr/testify/assert"
	"sauce-service/src/models"
	"gorm.io/gorm"
)

func cleanCategories(t *testing.T, db *gorm.DB) {
	err := db.Exec("DELETE FROM categories").Error
    if err != nil {
        t.Fatalf("failed to clean categories table: %v", err)
    }
}

func setupAndClean(t *testing.T) *gorm.DB {
    db := db.SetupTestDB(t)
    cleanCategories(t, db)
    return db
}

func createCategory(t *testing.T, db *gorm.DB, name string) *models.Category {
    cat, err := services.CreateCategory(db, name)
    if err != nil {
        t.Fatalf("failed to create category: %v", err)
    }
    return cat
}

func TestCreateCategory(t *testing.T) {
    t.Run("CreateCategory_Success", func(t *testing.T) {
		db := setupAndClean(t)
        cat := createCategory(t, db, "Sauces piquantes")
        assert.Equal(t, "Sauces piquantes", cat.Name)
    })

    t.Run("CreateCategory_EmptyName", func(t *testing.T) {
		db := setupAndClean(t)
        cat, err := services.CreateCategory(db, "")
        assert.NoError(t, err)
        assert.Equal(t, "", cat.Name)
    })
}

func TestGetAllCategories(t *testing.T) {
    t.Run("GetAllCategories_Success", func(t *testing.T) {
		db := setupAndClean(t)
        createCategory(t, db, "Sucrée")
		createCategory(t, db, "Salée")
		categories, _ := services.GetAllCategories(db)
		assert.Len(t, categories, 2)
		assert.Equal(t, "Salée", categories[0].Name)
		assert.Equal(t, "Sucrée", categories[1].Name)
    })

    t.Run("GetAllCategories_EmptyDB", func(t *testing.T) {
		db := setupAndClean(t)
		categories, err := services.GetAllCategories(db)
		assert.NoError(t, err)
		assert.Len(t, categories, 0)
    })
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("GetCategoryByID_Success", func(t *testing.T) {
		db := setupAndClean(t)
		cat := createCategory(t, db, "Végétarienne")
		found, err := services.GetCategoryByID(db, cat.ID.String())
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, "Végétarienne", found.Name)
	})

	t.Run("GetCategoryByID_NotFound", func(t *testing.T) {
		db := setupAndClean(t)
		_, err := services.GetCategoryByID(db, "id-invalide")
		assert.Error(t, err)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("UpdateCategory_Success", func(t *testing.T) {
		db := setupAndClean(t)
		cat := createCategory(t, db, "Ancien nom")
		updated, err := services.UpdateCategory(db, cat.ID.String(), "Nouveau nom")
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, "Nouveau nom", updated.Name)
	})

	t.Run("UpdateCategory_InvalidID", func(t *testing.T) {
		db := setupAndClean(t)
		updated, err := services.UpdateCategory(db, "inexistant-id", "Nom test")
		assert.Error(t, err)
		assert.Nil(t, updated)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("DeleteCategory_Success", func(t *testing.T) {
		db := setupAndClean(t)
		cat := createCategory(t, db, "À supprimer")
		err := services.DeleteCategory(db, cat.ID.String())
		assert.NoError(t, err)
		_, err = services.GetCategoryByID(db, cat.ID.String())
		assert.Error(t, err)
	})

	t.Run("DeleteCategory_InvalidID", func(t *testing.T) {
		db := setupAndClean(t)
		err := services.DeleteCategory(db, "inexistant-id")
		assert.Error(t, err)
	})
}