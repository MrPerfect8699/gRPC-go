package service

import (
	"errors"
	"grpc-employee/internal/model"
	"testing"
)

var ErrorEmployeeNotFound = errors.New("employee not found")

type fakeEmployeeRepository struct{}

func (f *fakeEmployeeRepository) GetEmployee(id int32) (*model.Employee, error) {
	if id != 1 {
		return nil, ErrorEmployeeNotFound
	}
	return &model.Employee{
		ID:         1,
		Name:       "John Doe",
		Department: "Engineering",
	}, nil
}

func (f *fakeEmployeeRepository) GetAllEmployees() ([]*model.Employee, error) {
	return []*model.Employee{
		{
			ID:         1,
			Name:       "John Doe",
			Department: "Engineering",
		},
	}, nil
}

func TestEmployeeService_GetEmployee(t *testing.T) {

	repo := &fakeEmployeeRepository{}

	svc := NewEmployeeService(repo)

	emp, err := svc.GetEmployee(1)
	if err != nil {
		t.Fatalf("Failed to get employee: %v", err)
	}
	if emp == nil {
		t.Fatal("Expected employee, got nil")
	}
	if emp.ID != 1 {
		t.Fatalf("Expected employee ID 1, got %d", emp.ID)
	}
	if emp.Name != "John Doe" {
		t.Fatalf("Expected employee name 'John Doe', got '%s'", emp.Name)
	}
}

func TestEmployeeService_GetEmployee_NotFound(t *testing.T) {
	repo := &fakeEmployeeRepository{}

	svc := NewEmployeeService(repo)

	emp, err := svc.GetEmployee(99)
	if err == nil {
		t.Fatalf("Expected error got nil")
	}
	if !errors.Is(err, ErrorEmployeeNotFound) {
		t.Fatalf("expected %+v, got %+v", ErrorEmployeeNotFound, err)
	}
	if emp != nil {
		t.Fatalf("Expected employee nil, but got %v", emp)
	}
}
