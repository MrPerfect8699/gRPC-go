package main

import (
	"log"
	"net"

	"grpc-employee/generated/grpc-employee/generated"
	"grpc-employee/internal/handler"
	"grpc-employee/internal/repository"
	"grpc-employee/internal/service"
	"grpc-employee/server/interceptors"

	"google.golang.org/grpc"
)

// type EmployeeServer struct {
// 	generated.UnimplementedEmployeeServiceServer
// }

// func (s *EmployeeServer) GetEmployee(ctx context.Context, req *generated.EmployeeRequest) (*generated.EmployeeResponse, error) {
// 	log.Println("Inside GetEmployee")
// 	employee := &generated.EmployeeResponse{
// 		Id:         req.GetId(),
// 		Name:       "John Doe",
// 		Department: "Engineering",
// 	}

// 	return employee, nil
// }

// func (s *EmployeeServer) DeleteEmployee(ctx context.Context, req *generated.EmployeeRequest) (*generated.EmployeeResponse, error) {
// 	employee := &generated.EmployeeResponse{
// 		Id:         req.GetId(),
// 		Name:       "Deleted",
// 		Department: "N/A",
// 	}

// 	return employee, nil
// }

// func (s *EmployeeServer) GetAllEmployees(req *generated.EmployeeRequest, stream grpc.ServerStreamingServer[generated.EmployeeResponse]) error {
// 	employees := []generated.EmployeeResponse{
// 		{Id: 1, Name: "John", Department: "IT"},
// 		{Id: 2, Name: "Alice", Department: "HR"},
// 		{Id: 3, Name: "Bob", Department: "Finance"},
// 	}

// 	for _, emp := range employees {
// 		if err := stream.Send(&emp); err != nil {
// 			return err
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	return nil
// }

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repository := repository.NewEmployeeRepository()

	svc := service.NewEmployeeService(repository)

	handler := handler.NewEmployeeHandler(svc)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AuthInterceptors),
	)

	generated.RegisterEmployeeServiceServer(grpcServer, handler)

	log.Println("gRPC Server started on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
