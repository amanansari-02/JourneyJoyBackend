package common

import (
	"Gin/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JsonResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	res := gin.H{
		"status":  statusCode,
		"message": message,
	}
	if data != nil {
		res["data"] = data
	}
	c.JSON(http.StatusOK, res)
}

func ErrorJsonResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": statusCode,
		"error":  message,
	})
}

func FindJsonResponse(c *gin.Context, searchField string, value string, data interface{}, statusCode int, errMsg string) {
	if err := config.DB.Where(searchField+" = ?", value).First(&data).Error; err == nil {
		ErrorJsonResponse(c, statusCode, errMsg)
		return
	} else if err != gorm.ErrRecordNotFound {
		ErrorJsonResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}
}
