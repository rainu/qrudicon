package image

import (
	"image"
	"image/color"
)

func MergeQrudicon(qr, identicon image.Image) image.Image {
	return insert(qr, removeAlpha(identicon))
}

//alpha von identicon entfernen
func removeAlpha(img image.Image) image.Image {
	newImage := image.NewRGBA(image.Rect(img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Max.X, img.Bounds().Max.Y))

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {

			if _, _, _, a := img.At(x, y).RGBA(); a == 0 {
				newImage.Set(x, y, color.White)
			} else {
				newImage.Set(x, y, img.At(x, y))
			}
		}
	}

	return newImage
}

func insert(outer, inner image.Image) image.Image {
	newImage := image.NewRGBA(image.Rect(outer.Bounds().Min.X, outer.Bounds().Min.Y, outer.Bounds().Max.X, outer.Bounds().Max.Y))
	startX := newImage.Bounds().Max.X/2 - (inner.Bounds().Max.X / 2) - 1
	endX := startX + inner.Bounds().Max.X
	startY := newImage.Bounds().Max.Y/2 - (inner.Bounds().Max.Y / 2) - 1
	endY := startY + inner.Bounds().Max.Y

	for x := 0; x < newImage.Bounds().Max.X; x++ {
		for y := 0; y < newImage.Bounds().Max.Y; y++ {

			//outer section
			if (x < startX || x >= endX) || (y < startY || y >= endY) {
				newImage.Set(x, y, outer.At(x, y))
			} else {
				//inner section
				innerX := x - startX
				innerY := y - startY

				newImage.Set(x, y, inner.At(innerX, innerY))
			}
		}
	}

	return newImage
}
