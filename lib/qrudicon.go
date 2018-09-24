package lib

import (
	"image"
	qrudiconImage "qrudicon/image"
)

func NewSimpleQrudicon(content string, size uint) image.Image {
	qrImage := qrudiconImage.NewQR(size, content)
	identiconImg := qrudiconImage.NewIdenticon(uint(size/4), content)
	mergedImage := qrudiconImage.MergeQrudicon(qrImage, identiconImg)

	return mergedImage
}
