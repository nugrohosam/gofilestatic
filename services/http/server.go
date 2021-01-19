package http

import (
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gosampleapi/services/http/exceptions"
	"github.com/spf13/viper"
)

// Routes ...
var Routes *gin.Engine

// Serve using for listen to specific port
func Serve() error {
	Prepare()

	port := viper.GetString("app.port")
	if err := Routes.Run(":" + port); err != nil {
		return err
	}

	return nil
}

// Prepare ...
func Prepare() {

	Routes = gin.New()

	isDebug := viper.GetBool("debug")
	if !isDebug {
		Routes.Use(exceptions.Recovery500())
	}

	rateLimiterCount := viper.GetUint64("rate-limiter.count")
	rateLimiterTime := viper.GetInt("rate-limiter.time-in-minutes")
	Routes.Use(throttle.Policy(&throttle.Quota{
		Limit:  rateLimiterCount,
		Within: time.Duration(rateLimiterTime) * time.Minute,
	}))

	Routes.Static("/assets", "./assets")
	Routes.Static("/web", "./web")

	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
}
