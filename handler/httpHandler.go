package handler

import (
	"ares/helper/response"
	"ares/helper/vo"
	"ares/proto"
	"ares/service"

	"github.com/gofiber/fiber/v3"
)

type HttpHandler struct {
	Service service.ServiceInt
}

func NewHttp(service service.ServiceInt) *HttpHandler {
	return &HttpHandler{
		Service: service,
	}
}

func (h *HttpHandler) HealthCheck(c fiber.Ctx) error {
	return c.JSON(response.SUCCESS(nil))
}

func (h *HttpHandler) Transaction(c fiber.Ctx) error {
	appCtx := vo.Parse(c)
	bind := c.Bind()

	request := proto.Request{}
	err := bind.Body(&request)
	if err != nil {
	}

	resp, err := h.Service.Transaction(appCtx.Session, &request)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
