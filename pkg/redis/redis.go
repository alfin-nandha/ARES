package redis

import (
	"ares/pkg/config"
	"ares/pkg/session"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInt interface {
	SetSession(session *session.Session)
	DeleteRedisKey(key string) error
	GetRedisKey(key string) (string, int)
	SaveRedis(key string, val any) error
	SaveRedisExp(key string, minutes string, val any) error
}

type Redis struct {
	session *session.Session
	client  *redis.Client
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
		client: redisClient,
	}
}

func (r *Redis) SetSession(s *session.Session) {
	r.session = s
}

/*
Redis Standard Get
*/
func (r *Redis) GetRedisKey(key string) (string, int) {

	val2, err := r.client.Get(r.session.Ctx, key).Result()
	if err == redis.Nil {
		r.session.LogInfo("REDIS Get Key Not Found", key)
		return val2, 1
	} else if err != nil {
		r.session.LogInfo("REDIS Get Key + Error", err.Error())
		return val2, -1

	} else {
		r.session.LogInfo("REDIS Get Key Success", key)
		return val2, 0
	}

}

/*
Redis Standard Set
*/
func (r *Redis) SaveRedis(key string, val any) error {
	var err error
	for i := 0; i < 3; i++ {
		err = r.client.Set(r.session.Ctx, key, val, 0).Err()
		if err == nil {
			r.session.LogInfo("REDIS save key success", key)
			return err
		}
	}
	r.session.LogError("REDIS Save Key "+key+" Error ", err.Error())
	return err
}

/*
Redis Standard Set Expired
*/
func (r *Redis) SaveRedisExp(key string, menit string, val any) error {
	var err error
	for i := 0; i < 3; i++ {
		duration, _ := time.ParseDuration(menit + "s")

		err := r.client.Set(r.session.Ctx, key, val, duration).Err()
		if err == nil {
			r.session.LogInfo("REDIS save key " + key + " success with duration " + duration.String())
			return err
		}
	}
	r.session.LogError("REDIS Save Key "+key+" Error ", err)
	return err
}

/*
Redis Standard Delete
*/
func (r *Redis) DeleteRedisKey(key string) error {
	var err error
	for i := 0; i < 3; i++ {
		err := r.client.Del(r.session.Ctx, key).Err()
		if err == nil {
			r.session.LogInfo("REDIS Delete key " + key + " success")
			break
		}
	}
	return err
}
