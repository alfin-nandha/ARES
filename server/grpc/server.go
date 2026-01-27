package grpc

import (
	"ares/container"
	"ares/pkg/config"
	"ares/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer(ctx context.Context, presenter *container.Presenter) (err error) {
	conf := config.Param.Apps

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		return
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			sessionInterceptor,
			panicRecoveryInterceptor,
			logInterceptor,
		),
	)
	if conf.Mode != "prod" {
		reflection.Register(srv)
	}

	proto.RegisterRuleEngineServer(srv, presenter.GrpcHandler)

	go func() {
		<-ctx.Done()
		log.Println("Shutting down gRPC...")
		srv.GracefulStop()
	}()

	err = srv.Serve(lis)
	return
}
