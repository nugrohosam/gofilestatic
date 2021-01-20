package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

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
func GetFileFromStorage(filepath string) ([]byte, error) {

	storageUse := viper.GetString("storage.driver")
	rootPathUse := viper.GetString("storage.root-path")
	var data []byte
	var err error

	switch storageUse {
	case StorageLocal:
		path := SetPath(rootPathUse, filepath)
		data, err = ioutil.ReadFile(path)
	case StorageGoogle:
	case StorageAws:
	}

	return data, err
}

// DuplicateToCache ..
func DuplicateToCache(fileToCopy []byte, filepath string) error {

	storageUse := viper.GetString("cache.driver")
	rootPathUse := viper.GetString("cache.root-path")
	var err error

	switch storageUse {
	case StorageLocal:
		path := SetPath(rootPathUse, filepath)
		err = ioutil.WriteFile(path, fileToCopy, os.ModeTemporary)
	case StorageGoogle:
	case StorageAws:
	}

	return err
}

// GetFileFromCache ..
func GetFileFromCache(filepath string) ([]byte, error) {

	storageUse := viper.GetString("cache.driver")
	rootPathUse := viper.GetString("cache.root-path")
	var data []byte
	var err error

	switch storageUse {
	case StorageLocal:
		path := SetPath(rootPathUse, filepath)
		data, err = ioutil.ReadFile(path)
	case StorageGoogle:
	case StorageAws:
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
