package sentry

import (
	"github.com/getsentry/sentry-go"
)

func SendErrorToSentry(err error) {
	sentry.WithScope(func(scope *sentry.Scope) {
		sentry.CaptureException(err)
	})
}
