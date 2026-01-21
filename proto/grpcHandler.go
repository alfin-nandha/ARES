package proto

import (
	"ares/service"
	context "context"
)

type Handler struct {
	UnimplementedRuleEngineServer
	service service.ServiceInt
}

func NewGrpcHandler(service service.ServiceInt) RuleEngineServer {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Transaction(context.Context, *Request) (*Response, error) {
	return nil, nil
}
