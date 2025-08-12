package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sauce-service/src/utils"
	app "sauce-service/src/app/stock"
	"sauce-service/src/services"
)

func CreateStock(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input app.CreateStockInput
		if !utils.BindAndValidateJSON(c, &input) {
			return
		}
		newStock, err := app.CreateStockFromInput(db, input)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondSuccess(c, http.StatusCreated, newStock)
	}
}

func GetAllStocks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		stocks, err := services.GetAllStocks(db)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondSuccess(c, http.StatusOK, stocks)
	}
}

func GetStockByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		stock, err := services.GetStockByID(db, id)
		if err != nil {
			utils.HandleError(c, err, "Stock not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, stock)
	}
}

func DeleteStock(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		if err := services.DeleteStock(db, id); err != nil {
			utils.HandleError(c, err, "Category not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Stock deleted successfully"})
	}
}

func UpdateStock(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := utils.ExtractIDParam(c)
		var input app.UpdateStockInput
		if !utils.BindAndValidateJSON(c, &input) {
			return
		}
		stock, err := app.UpdateStockFromInput(db, id, input)
		if err != nil {
			utils.HandleError(c, err, "Stock not found")
			return
		}
		utils.RespondSuccess(c, http.StatusOK, stock)
	}
}