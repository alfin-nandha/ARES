package server

import (
	"ares/helper/vo"
	"ares/pkg/session"
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
			stackTrace := string(debug.Stack()) // Capture stack trace
			log.Println(stackTrace)
			err = status.Errorf(codes.Internal, "an unexpected error occurred")
		}
	}()
	return handler(ctx, req)
}

func SessionInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata provided")
	}

	traceId := fmt.Sprintf("trc-%d", time.Now().UnixNano())
	metaTraceId := md.Get(vo.TraceId)
	if len(metaTraceId) != 0 {
		traceId = metaTraceId[0]
	}

	newSess := session.New().
		SetTraceId(traceId).
		SetRequest(req).
		SetURL(info.FullMethod).
		SetMetaData(md).
		SetMethod("gRPC")
	newCtx := context.WithValue(ctx, vo.AppSession, *newSess)

	return handler(newCtx, req)
}

func LogInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	start := time.Now()
	session := vo.Parse(ctx).Session
	session.LogRequest(nil)

	resp, err = handler(ctx, req)

	st, ok := status.FromError(err)
	var code codes.Code
	var msg string
	if ok {
		code = st.Code()
		msg = st.Message()
	} else {
		code = codes.Unknown
		msg = err.Error()
	}
	duration := time.Since(start)
	session.LogResponse(duration, code.String(), resp, msg)

	return
}

func AuthInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata provided")
	}

	traceId := fmt.Sprintf("trc-%d", time.Now().UnixNano())
	metaTraceId := md.Get(vo.TraceId)
	if len(metaTraceId) != 0 {
		traceId = metaTraceId[0]
	}

	newSess := session.New().
		SetTraceId(traceId).
		SetRequest(req).
		SetURL(info.FullMethod).
		SetMetaData(md).
		SetMethod("gRPC")
	newCtx := context.WithValue(ctx, vo.AppSession, *newSess)

	return handler(newCtx, req)
}
