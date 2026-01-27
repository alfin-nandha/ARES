package handler

import (
	"ares/helper/vo"
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
	appCtx := vo.Parse(ctx)

	resp, err = h.service.Transaction(appCtx.Session, req)
	return
}
