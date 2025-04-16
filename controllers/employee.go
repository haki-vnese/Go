package controllers

import (
	"go-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var employees = []models.Employee{}
var idCounter = 1

func CreateEmployee(c *gin.Context) {
	var newEmployee models.Employee
	if err := c.ShouldBindJSON(&newEmployee); err != nil { // ShouldBindJSON() read the value of the request body and assign it to newEmployee
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployee.ID = idCounter
	idCounter++
	employees = append(employees, newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func GetEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

func UpdateEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam) // convert string to int

	for i, emp := range employees { // loop through the employees slice
		if emp.ID == id {
			if err := c.ShouldBindJSON(&employees[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			employees[i].ID = id
			c.JSON(http.StatusOK, employees[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...) // append employees[i+1] -> end vao sau employees[:i] (i is skipped)
			c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}
