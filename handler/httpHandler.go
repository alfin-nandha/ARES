package handler

import (
	"ares/helper/response"
	"ares/proto"
	"ares/service"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/protobuf/encoding/protojson"
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
	request := proto.Request{}
	err := protojson.Unmarshal(c.Body(), &request)
	if err != nil {
	}

	resp, err := h.Service.Transaction(&request)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
