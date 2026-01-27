package server

import (
	"ares/container"
	"ares/server/grpc"
	"ares/server/http"
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

func StartServer() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	presenter, startUpSession := container.New()

	g.Go(func() error {
		startUpSession.LogInfo("Starting Grpc Server.....")
		return grpc.StartServer(ctx, presenter)
	})
	g.Go(func() error {
		startUpSession.LogInfo("Starting Fiber Server.....")
		return http.StartServer(ctx, presenter)
	})

	if err := g.Wait(); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
	log.Println("Everything closed safely.")
}
