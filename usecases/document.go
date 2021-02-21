package usecases

import (
	"path/filepath"
	"strings"

	"github.com/nugrohosam/gofilestatic/helpers"
	"github.com/nugrohosam/gofilestatic/services/infrastructure"
)

// GetDocumentInImage ..
func GetDocumentInImage(fileNameEncrypted string) (string, error) {

	return "", nil
}

// GetDocumentInPdf ..
func GetDocumentInPdf(fileNameEncrypted string) (string, error) {
	extWithDots := filepath.Ext(fileNameEncrypted)
	ext := strings.ReplaceAll(extWithDots, ".", "")
	if ext == "pdf" {
		return GetFile(fileNameEncrypted, "document")
	}

	var err error

	for {
		filePathCache := strings.ReplaceAll(fileNameEncrypted, extWithDots, "") + ".pdf"
		fileFullPathCache := helpers.CachePath(filePathCache)

		_, err = helpers.GetFileFromCache(filePathCache)
		if err == nil {
			return fileFullPathCache, nil
		}

		ext := filepath.Ext(fileNameEncrypted)
		secretImage := helpers.GetSecret("document", ext)
		pathFile := helpers.Decrypt(fileNameEncrypted, secretImage)
		storageFilePath := helpers.StoragePath(pathFile)
		folderCache := helpers.CachePath("")

		err = infrastructure.Convert("pdf", folderCache, storageFilePath)

		fileName := filepath.Base(pathFile)
		fileNameNewInPdf := strings.ReplaceAll(fileName, extWithDots, "") + ".pdf"
		filePathNowInCache := helpers.CachePath(fileNameNewInPdf)
		helpers.StorageMove(filePathNowInCache, fileFullPathCache)
	}
}
