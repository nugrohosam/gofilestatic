package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/nugrohosam/gofilestatic/helpers"
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

		fileData, err := ioutil.ReadFile(file.Filename)
		usecases.StoreImage(fileData, filePath, file.Filename)
	}
}

// ImageHandlerVerySmall is use
func ImageHandlerVerySmall() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
