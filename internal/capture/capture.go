package capture

import (
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func captureScreen() (image.Image, error) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)	
	if err != nil {
			return nil, err
		}
		return img, nil
}

func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}