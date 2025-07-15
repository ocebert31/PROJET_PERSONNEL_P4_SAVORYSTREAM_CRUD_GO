package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
)

func HandleError(c *gin.Context, err error, notFoundMsg string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		RespondError(c, http.StatusNotFound, notFoundMsg)
	} else {
		RespondError(c, http.StatusInternalServerError, err.Error())
	}
}

func ExtractIDParam(c *gin.Context) string {
	return c.Param("id")
}

func BindAndValidateJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

func RespondError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func RespondSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
