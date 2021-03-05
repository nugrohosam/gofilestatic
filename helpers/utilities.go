package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

const modeCopyFile = "755"

// Encrypt ..
func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err)
	}

	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

// Decrypt ..
func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

// GetFileFromStorage ..
func GetFileFromStorage(filePath string) ([]byte, error) {

	storageUse := viper.GetString("storage.driver")
	rootPathUse := viper.GetString("storage.root-path")
	var data []byte
	var err error

	switch storageUse {
	case STORAGE_LOCAL:
		path := SetPath(rootPathUse, filePath)
		data, err = ioutil.ReadFile(path)
	case STORAGE_GOOGLE:
	case STORAGE_AWS:
	}

	return data, err
}

// DuplicateToCache ..
func DuplicateToCache(fileToCopy []byte, filePath string) error {

	storageUse := viper.GetString("cache.driver")
	rootPathUse := viper.GetString("cache.root-path")
	var err error

	switch storageUse {
	case STORAGE_LOCAL:
		path := SetPath(rootPathUse, filePath)
		err = ioutil.WriteFile(path, fileToCopy, os.ModeTemporary)
	case STORAGE_GOOGLE:
	case STORAGE_AWS:
	}

	return err
}

// GetFileFromCache ..
func GetFileFromCache(filePath string) ([]byte, error) {

	storageUse := viper.GetString("cache.driver")
	rootPathUse := viper.GetString("cache.root-path")
	var data []byte
	var err error

	switch storageUse {
	case STORAGE_LOCAL:
		path := SetPath(rootPathUse, filePath)
		data, err = ioutil.ReadFile(path)
	case STORAGE_GOOGLE:
	case STORAGE_AWS:
	}

	return data, err
}

// SetPath ..
func SetPath(paths ...string) string {

	setPath := ""
	for _, path := range paths {
		setPath += "/" + path
	}

	return filepath.ToSlash(setPath)
}

// StoreImage ..
func StoreImage(file []byte, filePath string) error {

	storageUse := viper.GetString("image.driver")
	rootPathUse := viper.GetString("image.root-path")
	var err error

	switch storageUse {
	case STORAGE_LOCAL:
		path := SetPath(rootPathUse, filePath)
		err = ioutil.WriteFile(path, file, 0755)
	case STORAGE_GOOGLE:
	case STORAGE_AWS:
	}

	return err
}

// GetSecret ..
func GetSecret(typeFile, ext string) string {
	secretImage := viper.GetString("file-secret." + typeFile + "." + ext)
	if secretImage == "" {
		secretImage = viper.GetString("file-secret.other")
	}

	return secretImage
}

// MakeNameFile ..
func MakeNameFile(typeFile, ext string) string {
	secretFile := GetSecret("image", ext)
	uuidRandomString := uuid.MustParse(secretFile).String()
	return uuidRandomString + "." + ext
}

// ReadMultipartFile ..
func ReadMultipartFile(file *multipart.FileHeader) multipart.File {
	fileOpenned, err := file.Open()
	if err != nil {
		return fileOpenned
	}

	return fileOpenned
}
