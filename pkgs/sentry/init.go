package sentry

import (
	"api/configs"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("TZ", configs.ConfApp.AppTimezone)
	gin.SetMode(configs.ConfApp.AppMode)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfApp.SentryDNS,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	const sentryTimeoutSecond = 2 * time.Second

	sentry.Flush(sentryTimeoutSecond)
}

func SendErrorToSentry(err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		sentry.CaptureException(err)
	})
}
