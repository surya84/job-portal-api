package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig AppConfig
	DbConfig  DbConfig
	Redis     Redis
	Keys      Keys
}

type AppConfig struct {
	Port      string `env:"APP_PORT,required=true"`
	ReadTime  uint32 `env:"READ_TIME,required=true"`
	WriteTime uint32 `env:"WRITE_TIME,required=true"`
	Idle_Time uint32 `env:"IDLE_TIME,required=true"`
}

type DbConfig struct {
	DbConn string `env:"DB_DSN,required=true"`
}

type Redis struct {
	Host     string `env:"REDIS_HOST,default=localhost"`
	Port     string `env:"REDIS_PORT,default=6379"`
	Password string `env:"REDIS_PASSWORD,default=false"`
	DB       int    `env:"REDIS_DB,default=false"`
}
type Keys struct {
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
