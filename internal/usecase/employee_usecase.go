package usecase

import (
	"go-rest-api/internal/domain"
)

type EmployeeUsecase struct {
	repo domain.EmployeeRepository // EmployeeUsecase is a struct that implements the EmployeeUsecase interface
}

func NewEmployeeUsecase(repo domain.EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{
		repo: repo,
	}
}

func (u *EmployeeUsecase) Create(e *domain.Employee) error {
	return u.repo.Create(e)
}

func (u *EmployeeUsecase) GetByID(id int) (*domain.Employee, error) {
	return u.repo.GetByID(id)
}

func (u *EmployeeUsecase) GetAll() ([]domain.Employee, error) {
	return u.repo.GetAll()
}

func (u *EmployeeUsecase) Update(e *domain.Employee) error {
	return u.repo.Update(e)
}

func (u *EmployeeUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
