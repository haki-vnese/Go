package repository

import (
	"errors"
	"go-rest-api/internal/domain"
)

type EmployeeRepository struct { // EmployeeRepository is a struct that implements the EmployeeRepository interface
	db map[int]*domain.Employee // db is a map that simulates a database, where the key is the employee ID and the value is a pointer to an Employee struct
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		db: make(map[int]*domain.Employee),
	}
}

var idCounter = 0 // idCounter is a variable that simulates an auto-incrementing ID for new employees

func (r *EmployeeRepository) Create(e *domain.Employee) error {
	// Simulate creating an employee by adding it to the map
	e.ID = idCounter    // Set the ID of the employee to the current value of idCounter
	r.db[idCounter] = e // Add the employee to the map with the ID as the key
	idCounter++         // Increment the ID counter for the next employee
	return nil
}

func (r *EmployeeRepository) GetByID(id int) (*domain.Employee, error) {
	// Simulate getting an employee by ID by looking it up in the map
	employee, exist := r.db[id]
	if !exist {
		return nil, errors.New("employee not found") // Return an error if the employee is not found
	}
	return employee, nil // Return the employee if found
}

func (r *EmployeeRepository) GetAll() ([]domain.Employee, error) {
	employees := make([]domain.Employee, 0, len(r.db)) // create a temp array of Employee to avoid making changes to the original database
	for _, emp := range r.db {
		employees = append(employees, *emp)
	}
	return employees, nil // Return all employees as a slice
}

func (r *EmployeeRepository) Update(e *domain.Employee) error {
	// Simulate updating an employee by checking if it exists in the map and updating its values
	_, exist := r.db[e.ID]
	if !exist {
		return errors.New("employee not found") // Return an error if the employee is not found
	}
	r.db[e.ID] = e // Update the employee in the map
	return nil
}

func (r *EmployeeRepository) Delete(id int) error {
	// Simulate deleting an employee by checking if it exists in the map and deleting it
	_, exist := r.db[id]
	if !exist {
		return errors.New("employee not found") // Return an error if the employee is not found
	}
	delete(r.db, id) // Delete the employee from the map - I'm using a built-in function to delete a key-value pair from a map
	return nil
}
