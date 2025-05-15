package http

import (
	"go-rest-api/internal/domain"
	"go-rest-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeHandler is a struct that handles HTTP requests related to employees
type EmployeeHandler struct {
	usecase domain.EmployeeUsecase // usecase is an interface that defines the methods for employee use cases
}

// NewEmployeeHandler creates a new EmployeeHandler with the given employee use case
func NewEmployeeHandler(router *gin.Engine, uc domain.EmployeeUsecase) {
	h := &EmployeeHandler{usecase: uc}
	api := router.Group("/api")
	{
		api.POST("/employees", h.CreateEmployee)       // Route to create a new employee
		api.GET("/employees/:id", h.GetEmployeeByID)   // Route to get an employee by ID
		api.GET("/employees", h.GetAllEmployees)       // Route to get all employees
		api.PUT("/employees/:id", h.UpdateEmployee)    // Route to update an employee by ID
		api.DELETE("/employees/:id", h.DeleteEmployee) // Route to delete an employee by ID
	}
}

// CreateEmployee handles the HTTP request to create a new employee
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee domain.Employee                        // Create a new Employee struct to hold the request data
	if err := c.ShouldBindJSON(&employee); err != nil { // Bind the JSON request data to the Employee struct
		utils.ResponseError(c, http.StatusBadRequest, err.Error()) // Return a bad request error if binding fails
		return
	}

	if err := h.usecase.Create(&employee); err != nil { // Call the use case to create the employee
		utils.ResponseError(c, http.StatusInternalServerError, err.Error()) // Return an internal server error if creation fails
		return
	}

	utils.ResponseCreated(c, "Employee created", employee) // Return the created employee with a 201 status code
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")            // Get the employee ID from the URL parameter
	intID, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error()) // Return a bad request error if conversion fails
		return
	}
	employee, err := h.usecase.GetByID(intID) // Call the use case to get the employee by ID
	if err != nil {
		utils.ResponseError(c, http.StatusNotFound, err.Error()) // Return a not found error if the employee is not found
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Employee found", employee) // Return the employee with a 200 status code
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	employees, err := h.usecase.GetAll() // Call the use case to get all employees
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error()) // Return an internal server error if retrieval fails
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "All employees found", employees) // Return the list of employees with a 200 status code
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")            // Get the employee ID from the URL parameter
	intID, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error()) // Return a bad request error if conversion fails
		return
	}
	var employee domain.Employee                        // Create a new Employee struct to hold the request data
	if err := c.ShouldBindJSON(&employee); err != nil { // Bind the JSON request data to the Employee struct
		utils.ResponseError(c, http.StatusBadRequest, err.Error()) // Return a bad request error if binding fails
		return
	}
	employee.ID = intID // Set the ID of the employee to the ID from the URL parameter

	if err := h.usecase.Update(&employee); err != nil { // Call the use case to update the employee
		utils.ResponseError(c, http.StatusInternalServerError, err.Error()) // Return an internal server error if update fails
		return
	}
	utils.ResponseSuccess(c, http.StatusOK, "Employee updated", employee) // Return the updated employee with a 200 status code
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")            // Get the employee ID from the URL parameter
	intID, err := strconv.Atoi(id) // Convert the ID to an integer
	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error()) // Return a bad request error if conversion fails
		return
	}
	if err := h.usecase.Delete(intID); err != nil { // Call the use case to delete the employee by ID
		utils.ResponseError(c, http.StatusInternalServerError, err.Error()) // Return an internal server error if deletion fails
		return
	}
	c.Status(http.StatusNoContent) // Return a 204 No Content status code on successful deletion
}
