package infrastructure

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"
	"gopkg.in/gographics/imagick.v3/imagick"
)

// CreateImage ..
func CreateImage(quality uint, filePath, filePathDestination string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImage(filePath)
	mw.SetCompressionQuality(quality)
	fileDataCompressed := mw.GetImageBlob()

	folderOfFile := filepath.Dir(filePathDestination)
	helpers.FolderCheckAndCreate(folderOfFile)

	ioutil.WriteFile(filePathDestination, fileDataCompressed, 0744)
}

// CreateImageFromBlob ..
func CreateImageFromBlob(quality uint, sizePrecentage uint, file []byte, filePathDestination string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	mw.ReadImageBlob(file)

	sizedWidthFromPercentage := mw.GetImageWidth() * sizePrecentage / 100
	sizedHeightFromPercentage := mw.GetImageHeight() * sizePrecentage / 100

	fmt.Println(mw.GetImageWidth(), mw.GetImageHeight())
	fmt.Println(sizedWidthFromPercentage, sizedHeightFromPercentage)

	mw.SetCompressionQuality(quality)
	mw.SetSize(sizedWidthFromPercentage, sizedHeightFromPercentage)

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
