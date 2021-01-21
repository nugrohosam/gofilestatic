package usecases

import (
	"path/filepath"
	"strings"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/nugrohosam/gofilestatic/services/infrastructure"
	"github.com/spf13/viper"
)

// StoreImage ..
func StoreImage(file []byte, filePathRequsted, fileName string) error {
	ext := filepath.Ext(fileName)
	randomFileName := helpers.MakeNameFile(filePathRequsted, ext)
	filePath := helpers.SetPath(filePathRequsted, randomFileName)

	secretImage := helpers.GetSecret("image", ext)

	helpers.Encrypt(filePath, secretImage)
	return helpers.StoreImage(file, filePath)
}

// GetImageSmall ..
func GetImageSmall(fileNameEncrypted string) ([]byte, error) {

	prefixCachePath := viper.GetString("quality.image.very-small.prefix")

	var fileData []byte
	var err error
	fileData, err = helpers.GetFileFromCache(prefixCachePath + fileNameEncrypted)
	if err == nil {
		return fileData, nil
	}

	ext := filepath.Ext(fileNameEncrypted)
	filename := strings.ReplaceAll(fileNameEncrypted, ext, "")

	secretImage := helpers.GetSecret("image", ext)

	path := helpers.Decrypt(filename, secretImage)

	fileorgin := path + "." + ext
	filePathCache := prefixCachePath + fileNameEncrypted
	infrastructure.MakeImageVerySmall(fileorgin, filePathCache)

	return fileData, err
}
