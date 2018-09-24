package image

import (
	"github.com/skip2/go-qrcode"
	"image"
)

func NewQR(size uint, content string) image.Image {
	qr, err := qrcode.New(content, qrcode.Highest)
	if err != nil {
		panic(err)
	}

	return qr.Image(int(size))
}
