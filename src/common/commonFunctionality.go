package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
