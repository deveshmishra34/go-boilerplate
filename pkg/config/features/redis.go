package features

import "github.com/deveshmishra34/groot/pkg/utils/constants"

type RedisConfig struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

var redis = &Feature{
	Name:       constants.FEATURE_REDIS,
	Config:     &RedisConfig{},
	enabled:    true,
	configured: false,
	ready:      false,
	requirements: []string{
		"Host",
		"Port",
		"Password",
	},
}

func init() {
	Features.Add(redis)
}
