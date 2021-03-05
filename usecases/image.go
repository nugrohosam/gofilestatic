package usecases

import (
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

		switch quality {
		case IMAGE_VERY_SMALL_QUALITY:
			infrastructure.MakeImageVerySmall(fileData, fileFullPathCache)
		case IMAGE_SMALL_QUALITY:
			infrastructure.MakeImageSmall(fileData, fileFullPathCache)
		case IMAGE_MEDIUM_QUALITY:
			infrastructure.MakeImageMedium(fileData, fileFullPathCache)
		case IMAGE_LARGE_QUALITY:
			infrastructure.MakeImageLarge(fileData, fileFullPathCache)
		case IMAGE_VERY_LARGE_QUALITY:
			infrastructure.MakeImageVeryLarge(fileData, fileFullPathCache)
		default:
			return "", err
		}
	}
}
