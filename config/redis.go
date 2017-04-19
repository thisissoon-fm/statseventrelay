package config

import (
	"github.com/spf13/viper"
)

const (
	redis_host = "redis.host"
)

func init() {
	viper.SetDefault(redis_host, "localhost:6379")
}

func RedisHost() string {
	viper.BindEnv(redis_host)
	return viper.GetString(redis_host)
}
