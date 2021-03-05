package validations

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/thedevsaddam/govalidator"
)

// ShouldBeType ..
type ShouldBeType struct {
	Key      string
	Function func(string) govalidator.Options
}

// ValidateShouldBeString ...
var ValidateShouldBeString = ShouldBeType{
	Key:      "should-be-string",
	Function: ImageValidation,
}

// ImageValidation ..
func ImageValidation(attribute string) govalidator.Options {
	allowedFileExt := viper.GetStringSlice("rules.image.allowed-type")
	maxSize := viper.GetFloat64("rules.image.max")
	maxSizeInMB := int(1024 * 1024 * maxSize)
	ext := strings.Join(allowedFileExt, ",")

	fileAttribute := "file:" + attribute
	rules := govalidator.MapData{
		fileAttribute: []string{("ext:" + ext), ("size:" + strconv.Itoa(maxSizeInMB)), ("mime:" + ext), "required"},
	}

	messages := govalidator.MapData{
		fileAttribute: []string{("ext:Only " + ext + " is allowed"), "required:Photo is required"},
	}

	return govalidator.Options{
		Rules:    rules, // rules map,
		Messages: messages,
	}
}
