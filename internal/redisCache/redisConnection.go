package rediscache

import (
	"encoding/json"
	"fmt"
	"job-portal/config"
	"job-portal/internal/models"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func RedisClient() *redis.Client {
	cfg := config.GetConfig()
	//db, _ := strconv.Atoi(cfg.Redis.Db)
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),

		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return rdb
	//service.NewServiceRedis(rdb)

}

//go:generate mockgen -source=redisConnection.go -destination=redis_mock.go -package=rediscache
type Cache interface {
	SetRedisKey(key string, job models.Job)
	CheckRedisKey(key string) (models.Job, error)
	AddOtpToCache(email string, otp string)
	CheckOtpRequest(email string) (string, error)
	DeleteCacheData(email string) error
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

func (r *RdbConnection) AddOtpToCache(email string, otp string) {
	//otpNumber := strconv.Itoa(otp)
	err := r.rdb.Set(email, otp, 10*time.Minute).Err()
	if err != nil {
		log.Err(err)
		fmt.Println("Error storing OTP in Redis:", err)
		return
	}
	fmt.Println("OTP stored successfully in Redis for", email)
}

func (r *RdbConnection) CheckOtpRequest(email string) (string, error) {

	val, err := r.rdb.Get(email).Result()

	if err == redis.Nil {
		log.Err(err).Msg("Email data not found")
		return "", err
	}
	return val, nil

	//return val, err
	// if err == redis.Nil {

	// 	return false
	// }
	// return otp == val
}

func (r *RdbConnection) DeleteCacheData(email string) error {

	return r.rdb.Del(email).Err()
}
