package usecases

import (
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/nugrohosam/gofilestatic/services/infrastructure"
	"github.com/spf13/viper"
)

// LargeQuality ..
const LargeQuality = "large"

// VeryLargeQuality ..
const VeryLargeQuality = "very-large"

// MediumQuality ..
const MediumQuality = "medium"

// SmallQuality ..
const SmallQuality = "small"

// VerySmallQuality ..
const VerySmallQuality = "very-small"

// GetImageVerySmall ..
func GetImageVerySmall(fileNameEncrypted string) (string, error) {
	return getImage(VerySmallQuality, fileNameEncrypted)
}

// GetImageSmall ..
func GetImageSmall(fileNameEncrypted string) (string, error) {
	return getImage(SmallQuality, fileNameEncrypted)
}

// GetImageMedium ..
func GetImageMedium(fileNameEncrypted string) (string, error) {
	return getImage(MediumQuality, fileNameEncrypted)
}

// GetImageLarge ..
func GetImageLarge(fileNameEncrypted string) (string, error) {
	return getImage(LargeQuality, fileNameEncrypted)
}

// GetImageVeryLarge ..
func GetImageVeryLarge(fileNameEncrypted string) (string, error) {
	return getImage(VeryLargeQuality, fileNameEncrypted)
}

func getImage(quality, fileNameEncrypted string) (string, error) {

	cacheRootPath := viper.GetString("cache.root-path")
	prefixCachePath := viper.GetString("quality.image." + quality + ".prefix")
	var err error

	for {
		filePathCache := prefixCachePath + fileNameEncrypted
		fileFullPathCache := helpers.SetPath(cacheRootPath, filePathCache)

		_, err = helpers.GetFileFromCache(filePathCache)
		if err == nil {
			return fileFullPathCache, nil
		} else {
			err = nil
		}

		ext := filepath.Ext(fileNameEncrypted)
		secretImage := helpers.GetSecret("image", ext)
		pathFile := helpers.Decrypt(fileNameEncrypted, secretImage)
		fileData := helpers.GetFileDataStorage(pathFile)

		switch quality {
		case VerySmallQuality:
			infrastructure.MakeImageVerySmall(fileData, fileFullPathCache)
		case SmallQuality:
			infrastructure.MakeImageSmall(fileData, fileFullPathCache)
		case MediumQuality:
			infrastructure.MakeImageMedium(fileData, fileFullPathCache)
		case LargeQuality:
			infrastructure.MakeImageLarge(fileData, fileFullPathCache)
		case VeryLargeQuality:
			infrastructure.MakeImageVeryLarge(fileData, fileFullPathCache)
		default:
			return "", nil
		}
	}
}
