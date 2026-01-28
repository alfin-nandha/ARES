package model

import (
	"ares/proto"
)

type Trx struct {
	TransId string
	Payload map[string]any
}

func (t *Trx) Get(key string) any {
	return t.Payload[key]
}

func (t *Trx) Set(key string, value any) {
	t.Payload[key] = value
}

func (t *Trx) GetKeyContext() string {
	return "Trx"
}
func FromRequest(req *proto.Request) (trx Trx) {
	trx = Trx{
		TransId: req.TransactionId,
		Payload: req.Payload.AsMap(),
	}
	return
}
