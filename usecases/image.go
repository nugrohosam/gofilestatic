package usecases

import (
	"path/filepath"
	"strings"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/spf13/viper"
)

// GetImageSmall ..
func GetImageSmall(file string) ([]byte, error) {

	prefixCachePath := viper.GetString("quality.image.very-small.prefix")

	var fileData []byte
	var err error
	fileData, err = helpers.GetFileFromCache(prefixCachePath + file)
	if err == nil {
		return fileData, nil
	}

	ext := filepath.Ext(file)
	filename := strings.ReplaceAll(file, ext, "")

	secretImage := viper.GetString("file-secret.image." + ext)
	if secretImage == "" {
		secretImage = viper.GetString("file-secret.other")
	}

	path := helpers.Decrypt(filename, secretImage)

	fileorgin := path + "." + ext
	fileData, err = helpers.GetFileFromStorage(fileorgin)
	fileCache := prefixCachePath + file
	err = helpers.DuplicateToCache(fileData, fileCache)

	return fileData, err
}
