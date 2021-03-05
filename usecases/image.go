package usecases

import (
	"fmt"
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/nugrohosam/gofilestatic/services/infrastructure"
	"github.com/spf13/viper"
)

// GetImageSpecificQuality ..
func GetImageSpecificQuality(fileNameEncrypted, quality string) (string, error) {
	return get(quality, fileNameEncrypted)
}

func get(quality, fileNameEncrypted string) (string, error) {

	prefixCachePath := viper.GetString("quality.image." + quality + ".prefix")
	var err error

	for {
		filePathCache := prefixCachePath + fileNameEncrypted
		fileFullPathCache := helpers.CachePath(filePathCache)

		_, err = helpers.GetFileFromCache(filePathCache)
		if err == nil {
			return fileFullPathCache, nil
		}

		ext := filepath.Ext(fileNameEncrypted)
		secretImage := helpers.GetSecret("image", ext)
		pathFile := helpers.Decrypt(fileNameEncrypted, secretImage)
		fileData := helpers.GetFileDataStorage(pathFile)

		if helpers.InArray(quality, helpers.IMAGE_QUALITIES) {
			qualityCompression := viper.GetInt("quality.image." + quality + ".compression")
			sizePercentage := viper.GetInt("quality.image." + quality + ".size-percentage")
			fmt.Println(qualityCompression)
			fmt.Println(sizePercentage)
			infrastructure.CreateImageFromBlob(uint(qualityCompression), uint(sizePercentage), fileData, fileFullPathCache)
		} else {
			return "", err
		}
	}
}
