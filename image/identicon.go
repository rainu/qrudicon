package image

import (
	"github.com/jakobvarmose/go-qidenticon"
	"image"
)

func NewIdenticon(size uint, content string) image.Image {
	code := qidenticon.Code(content)
	settings := qidenticon.DefaultSettings()

	return qidenticon.Render(code, int(size), settings)
}
