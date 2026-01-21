package server

import (
	"ares/container"
	"ares/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			PanicRecoveryInterceptor,
		),
	)
	presenter := container.New()

	proto.RegisterRuleEngineServer(s, presenter.Handler)

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
