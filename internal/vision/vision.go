package vision

import "image"

func CropGameArea(img image.Image) image.Image {
	bounds := img.Bounds()

	rect := image.Rect(
		bounds.Min.X+100,
		bounds.Min.Y+100,
		bounds.Max.X-100,
		bounds.Max.Y-100,
	)

	if subImg, ok := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}); ok {
		return subImg.SubImage(rect)
	}
	return img
}