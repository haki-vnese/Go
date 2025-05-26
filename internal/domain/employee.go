package domain

// Interface for the Employee repo and usecase to communicate independently of the implementation
// This allows for easier testing and swapping of implementations if needed

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Position string `json:"position"`
}

type EmployeeRepository interface {
	Create(e *Employee) error
	GetByID(id int) (*Employee, error)
	GetAll() ([]Employee, error)
	Update(e *Employee) error
	Delete(id int) error
}

type EmployeeUsecase interface {
	Create(e *Employee) error
	GetByID(id int) (*Employee, error)
	GetAll() ([]Employee, error)
	Update(e *Employee) error
	Delete(id int) error
}
