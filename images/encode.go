package images

import (
	"image"
	"image/jpeg"
	"io"
)

func Encode(w io.Writer, img image.Image, imgDetail *Details) error {
	// Currently a wrapper for renderJpeg, but this function is useful if multiple render formats are needed
	return renderJpeg(w, img, imgDetail)
}

func renderJpeg(w io.Writer, m image.Image, imgDetail *Details) error {
	o := new(jpeg.Options)
	o.Quality = imgDetail.Quality

	return jpeg.Encode(w, m, o)
}