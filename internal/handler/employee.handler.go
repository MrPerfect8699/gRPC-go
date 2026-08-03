package handler

import (
	"context"
	"grpc-employee/generated/grpc-employee/generated"
	"grpc-employee/internal/port"
	"time"

	"google.golang.org/grpc"
)

type employeeHandler struct {
	service port.EmployeeService
	generated.UnimplementedEmployeeServiceServer
}

func NewEmployeeHandler(empService port.EmployeeService) *employeeHandler {
	return &employeeHandler{
		service: empService,
	}
}

func (h *employeeHandler) GetEmployee(ctx context.Context, req *generated.EmployeeRequest) (*generated.EmployeeResponse, error) {
	emp, err := h.service.GetEmployee(req.GetId())
	if err != nil {
		return nil, err
	}

	return &generated.EmployeeResponse{
		Id:         emp.ID,
		Name:       emp.Name,
		Department: emp.Department,
	}, nil
}

func (h *employeeHandler) GetAllEmployees(empReq *generated.EmployeeRequest, stream grpc.ServerStreamingServer[generated.EmployeeResponse]) error {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		return err
	}

	for _, emp := range employees {
		if err := stream.Send(&generated.EmployeeResponse{
			Id:         emp.ID,
			Name:       emp.Name,
			Department: emp.Department,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}
