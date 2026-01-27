package service

import (
	"ares/helper/vo"
	"ares/proto"
	context "context"
)

type Handler struct {
	proto.UnimplementedRuleEngineServer
	service ServiceInt
}

func NewGrpcHandler(service ServiceInt) proto.RuleEngineServer {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Transaction(ctx context.Context, req *proto.Request) (resp *proto.Response, err error) {
	appCtx := vo.Parse(ctx)

	resp, err = h.service.Transaction(appCtx.Session, req)
	return
}
