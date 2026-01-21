package server

import (
	"ares/pkg/session"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic and stack trace
			log.Printf("Panic captured in %s: %v", info.FullMethod, r)
			err = status.Errorf(codes.Internal, "an unexpected error occurred")
		}
	}()
	return handler(ctx, req)
}

func SessionInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	traceID := fmt.Sprintf("trc-%d", time.Now().UnixNano())

	newSess := session.New().SetTraceId(traceID)
	newSess.LogRequest()

	newCtx := context.WithValue(ctx, "sessionKey", newSess)

	return handler(newCtx, req)
}

func LogInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {

	log.Printf("REQ  Method: %s | Payload: %+v", info.FullMethod, req)

	start := time.Now()
	resp, err = handler(ctx, req)

	duration := time.Since(start)

	if err != nil {
		log.Printf("RESP Method: %s | Duration: %v | Error: %v", info.FullMethod, duration, err)
	} else {
		log.Printf("RESP Method: %s | Duration: %v | Payload: %+v", info.FullMethod, duration, resp)
	}

	return
}
