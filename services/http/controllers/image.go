package controllers

import (
	"net/http"

	"github.com/nugrohosam/gofilestatic/helpers"
	requests "github.com/nugrohosam/gofilestatic/services/http/requests/v1"
	"github.com/nugrohosam/gofilestatic/usecases"

	"github.com/gin-gonic/gin"
)

// ImageHandlerUpload is use
func ImageHandlerUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.FormFile("file")
		filePath, isOk := c.GetPostForm("filepath")
		if !isOk || err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}
		filePath := usecases.StoreFile(imageBuff, imageUpload.FilePath, imageUpload.File.Filename, "image")

		c.JSON(200, gin.H{
			"status":    true,
			"file_name": filePath,
		})
	}
}

// LargeQuality ..
const LargeQuality = "large"

// LargeQuality ..
const Original = "original"

// VeryLargeQuality ..
const VeryLargeQuality = "very-large"

// MediumQuality ..
const MediumQuality = "medium"

// SmallQuality ..
const SmallQuality = "small"

// VerySmallQuality ..
const VerySmallQuality = "very-small"

// ImageHandlerVerySmall is use
func ImageHandlerVerySmall() gin.HandlerFunc {
		file := c.Param("file")
		quality := c.Param("quality")

		var filePath string

		switch quality {
		case VerySmallQuality:
			filePath, _ = usecases.GetImageVerySmall(file)
		case SmallQuality:
			filePath, _ = usecases.GetImageSmall(file)
		case MediumQuality:
			filePath, _ = usecases.GetImageMedium(file)
		case LargeQuality:
			filePath, _ = usecases.GetImageLarge(file)
		case VeryLargeQuality:
			filePath, _ = usecases.GetImageVeryLarge(file)
		case Original:
			filePath, _ = usecases.GetFile(file, "image")
		default:
			c.Data(http.StatusNotFound, "Not Found", []byte("404 Not Found"))
			return
		}

		c.File(filePath)
	}
}
