package config

import (
	"github.com/spf13/viper"
)

const (
	rabbitmq_host = "rabbitmq.host"
	rabbitmq_user = "rabbitmq.user"
	rabbitmq_pass = "rabbitmq.pass"
	rabbitmq_port = "rabbitmq.port"
)

func init() {
	viper.SetDefault(rabbitmq_host, "localhost")
	viper.SetDefault(rabbitmq_user, "guest")
	viper.SetDefault(rabbitmq_pass, "guest")
	viper.SetDefault(rabbitmq_port, 5672)
}

func RabbitMQHost() string {
	viper.BindEnv(rabbitmq_host)
	return viper.GetString(rabbitmq_host)
}

func RabbitMQUser() string {
	viper.BindEnv(rabbitmq_user)
	return viper.GetString(rabbitmq_user)
}

func RabbitMQPass() string {
	viper.BindEnv(rabbitmq_pass)
	return viper.GetString(rabbitmq_pass)
}

func RabbitMQPort() int {
	viper.BindEnv(rabbitmq_port)
	return viper.GetInt(rabbitmq_port)
}
