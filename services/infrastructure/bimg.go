package infrastructure

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

// MakeVerySmall ..
func MakeVerySmall(filepath string) {
	options := bimg.Options{
		Width:     800,
		Height:    600,
		Crop:      true,
		Quality:   95,
		Rotate:    180,
		Interlace: true,
	}

	buffer, err := bimg.Read(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("new.jpg", newImage)
}
