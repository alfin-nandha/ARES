package container

import (
	"ares/handler"
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
	HttpHandler handler.HttpHandler
	Redis       redis.RedisInt
	Service     service.ServiceInt
	Repository  repository.RepositoryInt
}

func New() (*Presenter, *session.Session) {
	param := config.Param
	logger.New(param.Logger)

	startUpSession := session.New()
	startUpSession.SetTraceId("START UP")

	dbClient := database.New(startUpSession, param.Database)
	redis := redis.New(startUpSession, param.Cache)
	repo := repository.New(dbClient)
	serv := service.New(repo, redis)

	grpcHandler := handler.NewGrpc(serv)
	httpHandler := handler.NewHttp(serv)
	return &Presenter{
		GrpcHandler: grpcHandler,
		HttpHandler: *httpHandler,
		Service:     serv,
		Repository:  repo,
	}, startUpSession
}
