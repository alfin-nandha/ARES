package container

import (
	"ares/pkg/config"
	"ares/pkg/database"
	"ares/pkg/logger"
	"ares/pkg/redis"
	"ares/pkg/session"
	"ares/proto"
	"ares/repository"
	"ares/service"
)

type Presenter struct {
	GrpcHandler proto.RuleEngineServer
	Redis       redis.RedisInt
	service     service.ServiceInt
}

func New() Presenter {
	param := config.Param
	logger.New(param.Logger)

	startUpSession := session.New()
	startUpSession.SetTraceId("START UP")

	dbClient := database.New(startUpSession, param.Database)
	redis := redis.New(startUpSession, param.Cache)
	repo := repository.New(dbClient)
	serv := service.New(repo, redis)

	grpcHandler := service.NewGrpcHandler(serv)
	return Presenter{
		GrpcHandler: grpcHandler,
		service:     serv,
	}
}
