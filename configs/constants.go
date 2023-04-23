package configs

import "time"

const (
	DefaultDateCsvLayoutFormat = "02/01/2006"
	AccessTokenExpireTime      = 15 * time.Minute
	RefreshTokenExpireTime     = 24 * time.Hour
	ReadHeaderTimeout          = 5 * time.Second
)
