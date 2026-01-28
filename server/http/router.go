package http

import (
	Container "ares/container"
	"ares/pkg/config"

	"github.com/gofiber/fiber/v3"
)

func routerSetup(engine *fiber.App, presenter Container.Presenter) {
	app := engine.Group(config.Param.Apps.BaseUrl)

	api := app.Group("/api")
	api.Get("/health", presenter.HttpHandler.HealthCheck)

	v1 := api.Group("/v1")
	// v1.All("/*", presenter.Controller.Simualtor)
	v1.Post("/transaction", authMiddleware(&presenter), presenter.HttpHandler.Transaction)

}
