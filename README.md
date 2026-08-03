# gRPC Go Project

#grpc #golang #go #microservices #backend #protobuf #api #stackblitz #learning #cleanarchitecture

This project is a hands-on implementation of a gRPC service in Go using clean architecture principles. It demonstrates how to build a simple employee management service with separate layers for transport, business logic, repositories, and domain models.

## What this project includes

- A gRPC server and client written in Go
- Protocol buffer definitions for the employee service
- Clean architecture structure with handlers, services, ports, and repositories
- Basic error handling and unit tests for the service layer

## Project structure

- client/: gRPC client entry point
- server/: gRPC server and interceptors
- internal/handler/: request handling layer
- internal/service/: business logic
- internal/port/: interfaces for services and repositories
- internal/repository/: concrete repository implementation
- proto/: Protocol Buffers definitions
- generated/: generated gRPC and protobuf code

## Technologies used

- Go
- gRPC
- Protocol Buffers
- Clean architecture

## Purpose

This repository is intended for learning and practicing gRPC development in Go, as well as understanding how to organize a service using maintainable architectural patterns.
