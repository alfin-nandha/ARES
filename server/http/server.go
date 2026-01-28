package http

import (
	"ares/container"
	"ares/pkg/config"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func StartServer(ctx context.Context, presenter *container.Presenter) error {
	app := fiber.New()
	middlewareSetup(app, presenter)
	routerSetup(app, *presenter)

	go func() {
		<-ctx.Done() // Wait for Ctrl+C or errgroup failure
		log.Println("Shutting down Fiber...")
		_ = app.Shutdown()
	}()
	return app.Listen(fmt.Sprintf(":%d", config.Param.Apps.HttpPort))
}
