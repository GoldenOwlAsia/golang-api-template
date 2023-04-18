package configs

import (
	"github.com/spf13/viper"
	"time"
)

var (
	ConfApp Config
)

type Config struct {
	AppMode          string        `mapstructure:"APP_MODE"`
	AppTimezone      string        `mapstructure:"APP_TIMEZONE"`
	Port             string        `mapstructure:"PORT"`
	HttpSchema       string        `mapstructure:"HTTP_SCHEMA"`
	Domain           string        `mapstructure:"DOMAIN"`
	DatabaseUserName string        `mapstructure:"DATABASE_USER_NAME"`
	DatabasePassword string        `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost     string        `mapstructure:"DATABASE_HOST"`
	DatabaseName     string        `mapstructure:"DATABASE_NAME"`
	DatabasePort     int           `mapstructure:"DATABASE_PORT"`
	DatabaseTimezone string        `mapstructure:"DATABASE_TIMEZONE"`
	SentryDNS        string        `mapstructure:"SENTRY_DSN"`
	AllowOrigin      string        `mapstructure:"ALLOW_ORIGIN"`
	TokenSecret      string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn   time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge      int           `mapstructure:"TOKEN_MAX_AGE"`
	RedisEndpoint    string        `mapstructure:"REDIS_ENDPOINT"`
	RedisPassword    string        `mapstructure:"REDIS_PASSWORD"`
	MapsApiUrl       string        `mapstructure:"MAPS_API_URL"`
	MapsApiKey       string        `mapstructure:"MAPS_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	ConfApp = config
	return
}
