package usecases

import (
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"
)

// StoreFile ..
func StoreFile(file []byte, filePathRequsted, fileName, typeFile string) string {
	ext := filepath.Ext(fileName)
	randomFileName := helpers.MakeNameFile(filePathRequsted, ext)
	filePath := helpers.SetPath(filePathRequsted, randomFileName)

	secretImage := helpers.GetSecret(typeFile, ext)

	encrypted := helpers.Encrypt(filePath, secretImage)
	if err := helpers.StoreFile(file, filePath); err != nil {
		panic(err)
	}

	fileNameEncrypted := encrypted + ext

	return fileNameEncrypted
}

// GetFile ..
func GetFile(fileNameEncrypted, typeFile string) (string, error) {
	ext := filepath.Ext(fileNameEncrypted)
	secretImage := helpers.GetSecret(typeFile, ext)
	pathFile := helpers.Decrypt(fileNameEncrypted, secretImage)
	storageFilePath := helpers.StoragePath(pathFile)

	return storageFilePath, nil
}
