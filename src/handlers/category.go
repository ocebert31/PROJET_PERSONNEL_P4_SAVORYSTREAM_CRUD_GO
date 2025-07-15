package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sauce-service/src/utils"
	app "sauce-service/src/app/category"
	"sauce-service/src/services"
)

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input app.CategoryInput
		if !utils.BindAndValidateJSON(c, &input) {
			return
		}
		category, err := app.CreateFromInput(db, input)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondSuccess(c, http.StatusCreated, category)
	}
}

func GetAllCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := services.GetAllCategories(db)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondSuccess(c, http.StatusOK, categories)
	}
}

func GetCategoryByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		category, err := services.GetCategoryByID(db, id)
		if err != nil {
			utils.HandleError(c, err, "Category not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, category)
	}
}

func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		if err := services.DeleteCategory(db, id); err != nil {
			utils.HandleError(c, err, "Category not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Category deleted successfully"})
	}
}

func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		var input app.CategoryInput
		if !utils.BindAndValidateJSON(c, &input) {
			return
		}
		category, err := app.UpdateFromInput(db, id, input)
		if err != nil {
			utils.HandleError(c, err, "Category not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, category)
	}
}
