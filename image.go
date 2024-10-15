package ssd1322

import "github.com/fstanis/go-luma-ssd1322/pybridge"

// Image represents an image that can be drawn directly to the display.
type Image struct {
	img    *pybridge.Image
	width  int
	height int
	color  int
	mode   string
}

// NewImage creates a new image with the provided mode (such as "RGB"), width,
// height, and color.
func NewImage(mode string, width int, height int, color int) *Image {
	img := pybridge.ImageModule.NewImage(mode, width, height, color)
	if img == nil {
		return nil
	}
	return &Image{
		img:    img,
		width:  width,
		height: height,
		color:  color,
		mode:   mode,
	}
}

// Paste pastes the provided image on top of the current image.
func (i *Image) Paste(other *Image) {
	i.img.Paste(other.img)
}

// PasteXY pastes the provided image at the given coordinates on top of the
// current image.
func (i *Image) PasteXY(other *Image, x int, y int) {
	i.img.PasteXY(other.img, x, y)
}

// Crop returns a new image that is a cropped version of the current image.
func (i *Image) Crop(left int, top int, right int, bottom int) *Image {
	return &Image{
		img:    i.img.Crop(left, top, right, bottom),
		width:  right - left,
		height: bottom - top,
		color:  i.color,
		mode:   i.mode,
	}
}

// Convert returns a new image that is converted to the specified color mode.
func (i *Image) Convert(mode string) *Image {
	return &Image{
		img:    i.img.Convert(mode),
		width:  i.width,
		height: i.height,
		color:  i.color,
		mode:   mode,
	}
}

// Width returns the image width in pixels.
func (i *Image) Width() int {
	return i.width
}

// Height returns the image height in pixels.
func (i *Image) Height() int {
	return i.height
}

// Color returns the image color.
func (i *Image) Color() int {
	return i.color
}

// Mode returns the image color mode.
func (i *Image) Mode() string {
	return i.mode
}

// Free releases the resources associated with the image.
func (i *Image) Free() {
	i.img.Free()
}
