package service

import (
	"grpc-employee/internal/model"
	"grpc-employee/internal/port"
)

type employeeService struct {
	empService port.EmployeeRepository
}

func NewEmployeeService(empService port.EmployeeRepository) *employeeService {
	return &employeeService{
		empService: empService,
	}
}

func (s *employeeService) GetEmployee(id int32) (*model.Employee, error) {
	return s.empService.GetEmployee(id)
}

func (s *employeeService) GetAllEmployees() ([]*model.Employee, error) {
	return s.empService.GetAllEmployees()
}
