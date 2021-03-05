package infrastructure

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/spf13/viper"
	"gopkg.in/gographics/imagick.v3/imagick"
)

// MakeImageVerySmall ..
func MakeImageVerySmall(fileBlob []byte, filePathDestination string) {
	quality := viper.GetUint("quality.image.very-small.compression")
	CreateImageFormBlob(quality, fileBlob, filePathDestination)
}

// MakeImageSmall ..
func MakeImageSmall(fileBlob []byte, filePathDestination string) {
	quality := viper.GetUint("quality.image.small.compression")
	CreateImageFormBlob(quality, fileBlob, filePathDestination)
}

// MakeImageMedium ..
func MakeImageMedium(fileBlob []byte, filePathDestination string) {
	quality := viper.GetUint("quality.image.medium.compression")
	CreateImageFormBlob(quality, fileBlob, filePathDestination)
}

// MakeImageLarge ..
func MakeImageLarge(fileBlob []byte, filePathDestination string) {
	quality := viper.GetUint("quality.image.large.compression")
	CreateImageFormBlob(quality, fileBlob, filePathDestination)
}

// MakeImageVeryLarge ..
func MakeImageVeryLarge(fileBlob []byte, filePathDestination string) {
	quality := viper.GetUint("quality.image.very-large.compression")
	CreateImageFormBlob(quality, fileBlob, filePathDestination)
}

// CreateImage ..
func CreateImage(quality uint, filePath, filePathDestination string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImage(filePath)
	mw.SetCompressionQuality(quality)
	fileDataCompressed := mw.GetImageBlob()

	folderOfFile := filepath.Dir(filePathDestination)
	helpers.FolderCheckAndCreate(folderOfFile)

	ioutil.WriteFile(filePathDestination, fileDataCompressed, 0744)
}

// CreateImageFormBlob ..
func CreateImageFormBlob(quality uint, file []byte, filePathDestination string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImageBlob(file)
	mw.SetCompressionQuality(quality)
	fileDataCompressed := mw.GetImageBlob()

	folderOfFile := filepath.Dir(filePathDestination)
	helpers.FolderCheckAndCreate(folderOfFile)

	ioutil.WriteFile(filePathDestination, fileDataCompressed, 0744)
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
