package configs

import (
	"time"

	"github.com/spf13/viper"
)

var (
	ConfApp Config
)

type Config struct {
	AppMode          string        `mapstructure:"APP_MODE"`
	SecretKey        string        `mapstructure:"SECRET_KEY"`
	AppTimezone      string        `mapstructure:"APP_TIMEZONE"`
	Port             string        `mapstructure:"PORT"`
	HTTPSchema       string        `mapstructure:"HTTP_SCHEMA"`
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
}

func LoadConfig(path, nvFile string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(nvFile)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	ConfApp = config

	return
}
