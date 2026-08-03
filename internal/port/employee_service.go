package port

import "grpc-employee/internal/model"

type EmployeeService interface {
	GetAllEmployees() ([]*model.Employee, error)
	GetEmployee(id int32) (*model.Employee, error)
}
