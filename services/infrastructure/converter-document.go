package infrastructure

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

// Convert ..
func Convert(format, folderPathDestination, filePath string) error {
	apps := viper.GetString("converter.cmd")
	command := apps + " --headless --convert-to " + format + " --outdir " + folderPathDestination + " " + filePath
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()

	fmt.Println(string(output))
	return err
}
