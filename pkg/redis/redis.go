package redis

import (
	"ares/pkg/config"
	"ares/pkg/session"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInt interface {
	DeleteRedisKey(session *session.Session, key string) error
	GetRedisKey(session *session.Session, key string) (string, int)
	SaveRedis(session *session.Session, key string, val any) error
	SaveRedisExp(session *session.Session, key string, minutes string, val any) error
}

type Redis struct {
	Client *redis.Client
}

func New(session *session.Session, r config.Cache) RedisInt {
	redisClient := redis.NewClient(&redis.Options{
		Addr: r.Address,
	})

	_, errPing := redisClient.Ping(session.Ctx).Result()
	if errPing != nil {
		session.LogError("REDIS Not Connected", errPing.Error())
		log.Fatal("redis couldn't connect")
	}
	session.LogInfo("REDIS Connected")
	return &Redis{
		Client: redisClient,
	}
}

/*
Redis Standard Get
*/
func (r *Redis) GetRedisKey(session *session.Session, key string) (string, int) {

	val2, err := r.Client.Get(session.Ctx, key).Result()
	if err == redis.Nil {
		session.LogInfo("REDIS Get Key Not Found", key)
		return val2, 1
	} else if err != nil {
		session.LogInfo("REDIS Get Key + Error", err.Error())
		return val2, -1

	} else {
		session.LogInfo("REDIS Get Key Success", key)
		return val2, 0
	}

}

/*
Redis Standard Set
*/
func (r *Redis) SaveRedis(session *session.Session, key string, val any) error {
	var err error
	for i := 0; i < 3; i++ {
		err = r.Client.Set(session.Ctx, key, val, 0).Err()
		if err == nil {
			session.LogInfo("REDIS save key success", key)
			return err
		}
	}
	session.LogError("REDIS Save Key "+key+" Error ", err.Error())
	return err
}

/*
Redis Standard Set Expired
*/
func (r *Redis) SaveRedisExp(session *session.Session, key string, menit string, val any) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit + "s")

		err := r.Client.Set(session.Ctx, key, val, duration).Err()
		if err == nil {
			session.LogInfo("REDIS save key " + key + " success with duration " + duration.String())
			return err
		}
	}
	session.LogError("REDIS Save Key "+key+" Error ", err)
	return err
}

/*
Redis Standard Delete
*/
func (r *Redis) DeleteRedisKey(session *session.Session, key string) error {
	var err error
	for i := 0; i < 3; i++ {
		err := r.Client.Del(session.Ctx, key).Err()
		if err == nil {
			session.LogInfo("REDIS Delete key " + key + " success")
			break
		}
	}
	return err
}
