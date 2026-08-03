package repository

import (
	"grpc-employee/internal/model"
)

type employeeRepository struct {
	employees []*model.Employee
}

func NewEmployeeRepository() *employeeRepository {
	return &employeeRepository{
		employees: []*model.Employee{
			{ID: 1, Name: "John", Department: "IT"},
			{ID: 2, Name: "Alice", Department: "HR"},
			{ID: 3, Name: "Bob", Department: "Finance"},
		},
	}
}

// GetAllEmployees implements EmployeRepository.
func (e *employeeRepository) GetAllEmployees() ([]*model.Employee, error) {
	return e.employees, nil
}

// GetEmployee implements EmployeRepository.
func (e *employeeRepository) GetEmployee(id int32) (*model.Employee, error) {
	for _, emp := range e.employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return nil, EmployeeNotFoundError
}

// func (e *employeeRepository) DeleteEmployee(id int32) (*model.Employee, error) {
// 	for _, emp := range e.employees {
// 		if emp.ID == id {
// 			return emp, nil
// 		}
// 	}
// 	return nil, EmployeeNotFoundError
// }
