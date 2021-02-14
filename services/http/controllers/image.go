package controllers

import (
	"github.com/nugrohosam/gofilestatic/helpers"
	requests "github.com/nugrohosam/gofilestatic/services/http/requests/v1"

	"github.com/gin-gonic/gin"
)

// ImageHandlerUpload is use
func ImageHandlerUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var imageUpload requests.ImageUpload
		err := c.ShouldBind(&imageUpload)
		if err != nil {
			panic(err.Error())
		}

		_, err := helpers.ReadFileRequest(imageUpload.File)
	}
}

// ImageHandlerVerySmall is use
func ImageHandlerVerySmall() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
