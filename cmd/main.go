package main

import (
	"go-rest-api/internal/delivery/http"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Create a new Gin router

	// Initialize the repository and use case
	repo := repository.NewEmployeeRepository()
	uc := usecase.NewEmployeeUsecase(repo)

	// Initialize the HTTP handler with the router and use case
	http.NewEmployeeHandler(router, uc)

	// Start the server on port 8080
	router.Run(":8080")
}
