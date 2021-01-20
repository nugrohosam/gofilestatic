package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	httpConn "github.com/nugrohosam/gofilestatic/services/http"

	"github.com/nugrohosam/gofilestatic/services/infrastructure"
	"github.com/spf13/viper"
)

func main() {
	loadConfigFile()
	infrastructure.PrepareSentry()

	runHTTP := func() {
		if err := httpConn.Serve(); err != nil {
			panic(err)
		}
	}

	go runHTTP()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func loadConfigFile() {
	viper.SetConfigType("yaml")

	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Load all files in config folders
	var files []string

	configFolderName := "config"
	root := "./" + configFolderName
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() != configFolderName {
			files = append(files, info.Name())
		}
		return nil
	}); err != nil {
		panic(err)
	}

	var nameConfig string

	for _, file := range files {
		nameConfig = strings.ReplaceAll(file, ".yaml", "")

		viper.SetConfigName(nameConfig)
		viper.AddConfigPath(root)

		if err := viper.MergeInConfig(); err != nil {
			panic(err)
		}
	}
}
