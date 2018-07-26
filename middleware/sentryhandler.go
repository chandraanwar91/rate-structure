package middleware

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

type goMiddleware struct {
}

func (m *goMiddleware) SentryLogHandler() gin.HandlerFunc {
	return func(c *gin.Context){
		sentry.Recovery(raven.DefaultClient,false)
		c.Next()
	}
}

func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}