package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptors(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	log.Println("Interceptors Invoked")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No metadata received")
	}

	auth := md.Get("authorization")
	if len(auth) == 0 {
		return nil, status.Error(codes.Unauthenticated, "\n missing auth token")
	}

	if auth[0] != "Bearer my-token-123" {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}
