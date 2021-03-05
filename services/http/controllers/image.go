package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nugrohosam/gofilestatic/helpers"
	requests "github.com/nugrohosam/gofilestatic/services/http/requests/v1"
	"github.com/nugrohosam/gofilestatic/usecases"

	"github.com/gin-gonic/gin"
)

// ImageHandlerUpload is use
func ImageHandlerUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var imageUpload requests.ImageUpload
		c.Bind(&imageUpload)

		validate := helpers.NewValidation()
		if err := validate.Struct(imageUpload); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			fieldsErrors := helpers.TransformValidations(validationErrors)
			c.JSON(http.StatusBadRequest, helpers.ResponseErrValidation(fieldsErrors))
			return
		}

		path, err := helpers.SaveFileRequest(imageUpload.File)
		if err != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseErr(err.Error()))
			return
		}

		imageBuff, err := ioutil.ReadFile(path)
		if err != nil {
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

// ImageHandler is use
func ImageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		quality := c.Param("quality")

		var filePath string

		if helpers.InArray(quality, helpers.IMAGE_QUALITIES) {
			if quality == helpers.IMAGE_ORIGINAL {
				filePath, _ = usecases.GetFile(file, "image")
			} else {
				filePath, _ = usecases.GetImageSpecificQuality(file, quality)
			}
		} else {
			c.Data(http.StatusNotFound, "Not Found", []byte("404 Not Found"))
			return
		}

		c.File(filePath)
	}
}
