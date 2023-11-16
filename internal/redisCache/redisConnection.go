package rediscache

import (
	"encoding/json"
	"job-portal/internal/models"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
	//service.NewServiceRedis(rdb)

}

//go:generate mockgen -source=redisConnection.go -destination=redis_mock.go -package=rediscache
type Cache interface {
	SetRedisKey(key string, job models.Job)
	CheckRedisKey(key string) (models.Job, error)
}

type RdbConnection struct {
	rdb *redis.Client
}

func NewredisConnection(rdb *redis.Client) Cache {
	return &RdbConnection{
		rdb: rdb,
	}
}

func (r *RdbConnection) SetRedisKey(key string, job models.Job) {
	jobdata, err := json.Marshal(job)
	if err != nil {
		log.Err(err)
		return
	}
	data := string(jobdata)
	err = r.rdb.Set(key, data, 30*time.Minute).Err()
	if err != nil {
		log.Err(err)
		return
	}
}

func (r *RdbConnection) CheckRedisKey(key string) (models.Job, error) {

	val, err := r.rdb.Get(key).Result()
	if err == redis.Nil {
		return models.Job{}, err
	}
	var job models.Job
	err = json.Unmarshal([]byte(val), &job)
	if err != nil {
		log.Err(err)
	}
	return job, nil
}
