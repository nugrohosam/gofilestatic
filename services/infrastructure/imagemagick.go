package infrastructure

import (
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/gographics/imagick.v3/imagick"
)

// MakeImageVerySmall ..
func MakeImageVerySmall(filepath, filepathDestination string) {
	quality := viper.GetUint("quality.image.very-small.compression")
	CreateImage(quality, filepath, filepathDestination)
}

// MakeImageSmall ..
func MakeImageSmall(filepath, filepathDestination string) {
	quality := viper.GetUint("quality.image.small.compression")
	CreateImage(quality, filepath, filepathDestination)
}

// MakeImageMedium ..
func MakeImageMedium(filepath, filepathDestination string) {
	quality := viper.GetUint("quality.image.medium.compression")
	CreateImage(quality, filepath, filepathDestination)
}

// CreateImage ..
func CreateImage(quality uint, filepath, filepathDestination string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImage(filepath)
	mw.SetCompressionQuality(quality)
	fileDataCompressed := mw.GetImageBlob()

	ioutil.WriteFile(filepathDestination, fileDataCompressed, 0744)
}

// ConvertImage ..
func ConvertImage(imageType, filePath string) error {
	fileTmp, err := ioutil.TempFile("", filePath)
	if err != nil {
		return err
	}

	defer os.Remove(fileTmp.Name())
	fileDataCompressed := fileTmp.Name() + "." + imageType
	_, err = imagick.ConvertImageCommand([]string{"convert", fileTmp.Name(), fileDataCompressed})

	return err
}
