package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

func ResponseError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}

func ResponseSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": message,
		"data":    data,
	})
}
