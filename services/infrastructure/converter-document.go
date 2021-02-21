package infrastructure

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/nugrohosam/gofilestatic/helpers"

	"github.com/spf13/viper"
)

// CovertDocument ..
func CovertDocument(format, folderPathDestination, filePath string) error {

	librelistFormatAllowed := viper.GetStringSlice("converter.libre.alowed-format")
	vipslistFormatAllowed := viper.GetStringSlice("converter.vips.alowed-format")

	if helpers.InArray(format, vipslistFormatAllowed) {
		fileName := filepath.Base(filePath)
		fileDestinationPath := helpers.SetPath(folderPathDestination, fileName)

		apps := viper.GetString("converter.vips.cmd")
		command := apps + " " + filePath + " " + fileDestinationPath
		cmd := exec.Command("bash", "-c", command)
		output, err := cmd.Output()

		fmt.Println(string(output))
		return err
	} else if helpers.InArray(format, librelistFormatAllowed) {
		apps := viper.GetString("converter.libre.cmd")
		command := apps + " --headless --convert-to " + format + " --outdir " + folderPathDestination + " " + filePath
		cmd := exec.Command("bash", "-c", command)
		output, err := cmd.Output()

		fmt.Println(string(output))
		return err
	}

	return errors.New("Format not supported")
}
