package services_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sauce-service/src/models"
	"gorm.io/gorm"
	"sauce-service/src/utils"
	"sauce-service/src/services"
)

func createCategory(t *testing.T, db *gorm.DB, name string) *models.Category {
	cat, err := services.CreateCategory(db, name) 
	if err != nil {
		t.Fatalf("failed to create category: %v", err)
	}
	return cat
}

func TestCreateCategory(t *testing.T) {
    t.Run("CreateCategory_Success", func(t *testing.T) {
		utils.SetupTestDB()
		cat := createCategory(t, utils.TestDB, "Sauces piquantes")
        assert.Equal(t, "Sauces piquantes", cat.Name)
    })

    t.Run("CreateCategory_EmptyName", func(t *testing.T) {
		utils.SetupTestDB()
        cat, err := services.CreateCategory(utils.TestDB, "")
        assert.NoError(t, err)
        assert.Equal(t, "", cat.Name)
    })
}

func TestGetAllCategories(t *testing.T) {
    t.Run("GetAllCategories_Success", func(t *testing.T) {
		utils.SetupTestDB()
        createCategory(t, utils.TestDB, "Sucrée")
		createCategory(t, utils.TestDB, "Salée")
		categories, _ := services.GetAllCategories(utils.TestDB)
		assert.Len(t, categories, 2)
		assert.Equal(t, "Salée", categories[0].Name)
		assert.Equal(t, "Sucrée", categories[1].Name)
    })

    t.Run("GetAllCategories_EmptyDB", func(t *testing.T) {
		utils.SetupTestDB()
		categories, err := services.GetAllCategories(utils.TestDB)
		assert.NoError(t, err)
		assert.Len(t, categories, 0)
    })
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("GetCategoryByID_Success", func(t *testing.T) {
		utils.SetupTestDB()
		cat := createCategory(t, utils.TestDB, "Végétarienne")
		found, err := services.GetCategoryByID(utils.TestDB, cat.ID.String())
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, "Végétarienne", found.Name)
	})

	t.Run("GetCategoryByID_NotFound", func(t *testing.T) {
		utils.SetupTestDB()
		_, err := services.GetCategoryByID(utils.TestDB, "id-invalide")
		assert.Error(t, err)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("UpdateCategory_Success", func(t *testing.T) {
		utils.SetupTestDB()
		cat := createCategory(t, utils.TestDB, "Ancien nom")
		updated, err := services.UpdateCategory(utils.TestDB, cat.ID.String(), "Nouveau nom")
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, "Nouveau nom", updated.Name)
	})

	t.Run("UpdateCategory_InvalidID", func(t *testing.T) {
		utils.SetupTestDB()
		updated, err := services.UpdateCategory(utils.TestDB, "inexistant-id", "Nom test")
		assert.Error(t, err)
		assert.Nil(t, updated)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("DeleteCategory_Success", func(t *testing.T) {
		utils.SetupTestDB()
		cat := createCategory(t, utils.TestDB, "À supprimer")
		err := services.DeleteCategory(utils.TestDB, cat.ID.String())
		assert.NoError(t, err)
		_, err = services.GetCategoryByID(utils.TestDB, cat.ID.String())
		assert.Error(t, err)
	})

	t.Run("DeleteCategory_InvalidID", func(t *testing.T) {
		utils.SetupTestDB()
		err := services.DeleteCategory(utils.TestDB, "inexistant-id")
		assert.Error(t, err)
	})
}