package core

import (
	"ares/pkg/redis"
	"ares/proto"
	"ares/repository"
)

type Trx struct {
	TransId string
	Payload map[string]any
	redis   redis.RedisInt
	repo    repository.RepositoryInt
}

func New(req *proto.Request, redis redis.RedisInt, repo repository.RepositoryInt) (trx Trx) {
	trx = Trx{
		TransId: req.TransactionId,
		Payload: req.Payload.AsMap(),

		redis: redis,
		repo:  repo,
	}
	return
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

func (t *Trx) GetData(key string) string {
	str, _ := t.redis.GetRedisKey(key)
	return str
}
