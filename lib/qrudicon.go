package lib

import (
	"image"
	qrudiconImage "qrudicon/image"
)

func NewSimpleQrudicon(content string, size uint) image.Image {
	return NewExtendedQrudicon(content, content, size)
}

func NewExtendedQrudicon(qrContent, idContent string, size uint) image.Image {
	qrImage := qrudiconImage.NewQR(size, qrContent)
	identiconImg := qrudiconImage.NewIdenticon(uint(size/4), idContent)
	mergedImage := qrudiconImage.MergeQrudicon(qrImage, identiconImg)

	return mergedImage
}
