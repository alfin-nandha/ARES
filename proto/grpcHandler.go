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

func (h *Handler) Transaction(ctx context.Context, req *Request) (resp *Response, err error) {

	return nil, nil
}
