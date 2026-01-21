package container

import (
	"ares/pkg/config"
	"ares/pkg/database"
	"ares/pkg/logger"
	"ares/pkg/redis"
	"ares/proto"
	"ares/repository"
	"ares/service"
	"encoding/json"
	"fmt"
)

type Presenter struct {
	Handler proto.RuleEngineServer
	Redis   redis.RedisInt
	service service.ServiceInt
}

func New() Presenter {
	param := config.Param
	b, _ := json.Marshal(param)
	fmt.Println(string(b))
	logger.New(param.Logger)
	dbClient := database.New(param.Database)
	redis := redis.New(param.Cache)
	repo := repository.New(dbClient)
	serv := service.New(repo, redis)
	hand := proto.NewGrpcHandler(serv)
	return Presenter{
		Handler: hand,
		service: serv,
	}
}
