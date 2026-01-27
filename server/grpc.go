package server

import (
	"ares/container"
	"ares/pkg/config"
	"ares/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer() {
	conf := config.Param.Apps

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GrpcPort))
	if err != nil {
		log.Fatal(err)
		return
	}

	presenter := container.New()
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			SessionInterceptor,
			PanicRecoveryInterceptor,
			LogInterceptor,
		),
	)
	if conf.Mode != "prod" {
		reflection.Register(s)
	}

	proto.RegisterRuleEngineServer(s, presenter.GrpcHandler)

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
