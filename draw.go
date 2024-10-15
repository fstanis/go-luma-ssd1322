package ssd1322

import "github.com/fstanis/go-luma-ssd1322/pybridge"

// Draw represents a drawing context on top of an Image.
type Draw struct {
	imageDraw *pybridge.ImageDraw
}

// NewDraw creates a new drawing context on top of the provided Image.
func NewDraw(image *Image) *Draw {
	return &Draw{
		pybridge.ImageDrawModule.Draw(image.img),
	}
}

// Text draws the provided text at the given coordinates using the provided font.
func (i *Draw) Text(x int, y int, text string, font *Font) {
	i.imageDraw.Text(x, y, text, 255, font.font)
}

// Bitmap draws the provided bitmap at the given coordinates.
func (i *Draw) Bitmap(x int, y int, bitmap *Image) {
	i.imageDraw.Bitmap(x, y, bitmap.img, 255)
}

// Rectangle draws a rectangle at the given coordinates.
func (i *Draw) Rectangle(x1 int, y1 int, x2 int, y2 int, color int) {
	i.imageDraw.Rectangle(x1, y1, x2, y2, color)
}

// Free releases the resources associated with the drawing context.
func (d *Draw) Free() {
	d.imageDraw.Free()
}
