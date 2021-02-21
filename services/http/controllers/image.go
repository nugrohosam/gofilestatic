package controllers

import (
	"github.com/nugrohosam/gofilestatic/helpers"
	requests "github.com/nugrohosam/gofilestatic/services/http/requests/v1"
	"github.com/nugrohosam/gofilestatic/usecases"

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

		imageBuff, err := helpers.ReadFileRequest(imageUpload.File)
		filePath := usecases.StoreImage(imageBuff, imageUpload.FilePath, imageUpload.File.Filename)

		c.JSON(200, gin.H{
			"status":    true,
			"file_name": filePath,
		})
	}
}

// ImageHandlerVerySmall is use
func ImageHandlerVerySmall() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, err := usecases.GetImageVerySmall(file)
		if err != nil {
			panic(err)
		}

		c.File(filePath)
	}
}

// ImageHandlerSmall is use
func ImageHandlerSmall() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, err := usecases.GetImageSmall(file)
		if err != nil {
			panic(err)
		}

		c.File(filePath)
	}
}

// ImageHandlerMedium is use
func ImageHandlerMedium() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, err := usecases.GetImageMedium(file)
		if err != nil {
			panic(err)
		}

		c.File(filePath)
	}
}

// ImageHandlerLarge is use
func ImageHandlerLarge() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, err := usecases.GetImageLarge(file)
		if err != nil {
			panic(err)
		}

		c.File(filePath)
	}
}

// ImageHandlerVeryLarge is use
func ImageHandlerVeryLarge() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		filePath, err := usecases.GetImageVeryLarge(file)
		if err != nil {
			panic(err)
		}

		c.File(filePath)
	}
}
