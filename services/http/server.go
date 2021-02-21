package http

import (
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
	"github.com/nugrohosam/gofilestatic/services/http/controllers"
	"github.com/nugrohosam/gofilestatic/services/http/exceptions"
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

	Routes.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	v1 := Routes.Group("v1")

	image := v1.Group("image")
	{
		image.POST("", controllers.ImageHandlerUpload())
		image.GET("very-small/:file", controllers.ImageHandlerVerySmall())
		image.GET("small/:file", controllers.ImageHandlerSmall())
		image.GET("medium/:file", controllers.ImageHandlerMedium())
		image.GET("large/:file", controllers.ImageHandlerLarge())
		image.GET("very-large/:file", controllers.ImageHandlerVeryLarge())
	}
}
