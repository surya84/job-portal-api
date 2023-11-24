package rediscache

import (
	"encoding/json"
	"fmt"
	"job-portal/config"
	"job-portal/internal/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func RedisClient() *redis.Client {
	cfg := config.GetConfig()
	db, _ := strconv.Atoi(cfg.Redis.Db)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(":%s", cfg.Redis.Address),
		Password: cfg.Redis.Password,
		DB:       db,
	})

	return rdb
	//service.NewServiceRedis(rdb)

}

//go:generate mockgen -source=redisConnection.go -destination=redis_mock.go -package=rediscache
type Cache interface {
	SetRedisKey(key string, job models.Job)
	CheckRedisKey(key string) (models.Job, error)
	AddOtpToCache(email string, otp int)
	CheckOtpRequest(email string, otp string) bool
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

func (r *RdbConnection) AddOtpToCache(email string, otp int) {
	otpNumber := string(otp)
	err := r.rdb.Set(email, otpNumber, 10*time.Minute).Err()
	if err != nil {
		log.Err(err)
		return
	}
}

func (r *RdbConnection) CheckOtpRequest(email string, otp string) bool {

	val, err := r.rdb.Get(email).Result()

	if err != nil {
		return false
	}

	if otp == val {
		return true
	}
	return false
}
