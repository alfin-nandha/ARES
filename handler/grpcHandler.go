package handler

import (
	"ares/proto"
	"ares/service"
	context "context"
)

type GrpcHandler struct {
	proto.UnimplementedRuleEngineServer
	service service.ServiceInt
}

func NewGrpc(service service.ServiceInt) proto.RuleEngineServer {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) Transaction(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	resp, err = h.service.Transaction(req)
	return
}
