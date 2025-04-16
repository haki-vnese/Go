package controllers

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request data (email, password) through the database
	for _, user := range models.Users {
		if user.Email == req.Email && user.Password == req.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}
