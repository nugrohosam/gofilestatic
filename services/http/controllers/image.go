package controllers

import (
	"io/ioutil"
	"os"

	"github.com/nugrohosam/gofilestatic/usecases"

	"github.com/adelowo/filer/validator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ImageHandlerUpload is use
func ImageHandlerUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			return
		}

		filePath, isOk := c.GetPostForm("filepath")
		if isOk {
			return
		}

		fileTemp, err := ioutil.TempFile("", file.Filename)
		defer os.Remove(fileTemp.Name())

		if err != nil {
			panic("An error occurred while trying to create a temporary file")
		}

		max := viper.GetFloat64("rules.image.max")
		min := viper.GetFloat64("rules.image.min")
		allowedTypes := viper.GetStringSlice("rules.image.allowed-type")

		validateMime := validator.NewMimeTypeValidator(allowedTypes)
		validateSize := validator.NewSizeValidator(int64(1024*1024*max), int64(1024*1024*min)) //2MB(maxSize) and 200KB(minSize)

		if _, err := validateMime.Validate(fileTemp); err != nil {
			c.JSON(400, nil)
			return
		}

		if _, err := validateSize.Validate(fileTemp); err != nil {
			c.JSON(400, nil)
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
