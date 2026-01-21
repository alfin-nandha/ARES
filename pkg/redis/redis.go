package redis

import (
	"ares/pkg/config"
	"ares/pkg/logger"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInt interface {
	DeleteRedisKey(ctx context.Context, key string) error
	GetRedisKey(ctx context.Context, key string) (string, int)
	SaveRedis(ctx context.Context, key string, val interface{}) error
	SaveRedisExp(ctx context.Context, key string, menit string, val interface{}) error
}

type Redis struct {
	Client *redis.Client
}

func New(r config.Cache) RedisInt {
	redisClient := redis.NewClient(&redis.Options{
		Addr: r.Address,
	})
	ctx := context.Background()

	_, errPing := redisClient.Ping(ctx).Result()
	if errPing != nil {
		logger.Log.Error("REDIS Not Connected : " + errPing.Error())
		log.Fatal("redis couldn't connect")
	}
	logger.Log.Info("REDIS Connected")
	return &Redis{
		Client: redisClient,
	}
}

/*
Redis Standard Get
*/
func (r *Redis) GetRedisKey(ctx context.Context, key string) (string, int) {

	val2, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		logger.Log.Info("REDIS Get Key " + key + " Key Not Found")
		return val2, 1
	} else if err != nil {
		logger.Log.Error("REDIS Get Key " + key + " Error " + fmt.Sprintf("%v", err))
		return val2, -1

	} else {
		logger.Log.Info("REDIS Get Key " + key + " Success")
		return val2, 0

	}

}

/*
Redis Standard Set
*/
func (r *Redis) SaveRedis(ctx context.Context, key string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		err = r.Client.Set(ctx, key, val, 0).Err()
		if err == nil {
			logger.Log.Info("REDIS save key " + key + " success")
			return err
		}
	}
	logger.Log.Error("REDIS Save Key " + key + " Error " + fmt.Sprintf("%v", err))
	return err
}

/*
Redis Standard Set Expired
*/
func (r *Redis) SaveRedisExp(ctx context.Context, key string, menit string, val interface{}) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit + "s")

		err := r.Client.Set(ctx, key, val, duration).Err()
		if err == nil {
			logger.Log.Info("REDIS save key " + key + " success with duration " + duration.String())
			return err
		}
	}
	logger.Log.Error("REDIS Save Key " + key + " Error " + fmt.Sprintf("%v", err))
	return err
}

/*
Redis Standard Delete
*/
func (r *Redis) DeleteRedisKey(ctx context.Context, key string) error {
	var err error
	for i := 0; i < 3; i++ {
		err := r.Client.Del(ctx, key).Err()
		if err == nil {
			logger.Log.Info("REDIS Delete key " + key + " success")

			break
		}
	}
	return err
}
