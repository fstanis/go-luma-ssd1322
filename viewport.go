package ssd1322

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrImageSizeMismatch = errors.New("image size mismatch")
	ErrInvalidHotspot    = errors.New("invalid hotspot")
)

type widthHeight interface {
	Width() int
	Height() int
}

type hotspotable interface {
	widthHeight
	ShouldRedraw() bool
	PasteInto(image *Image, x, y int)
}

type hotspotWithXY struct {
	hotspot hotspotable
	x       int
	y       int
}

// Viewport is an interface to a device that allows adding multiple "hotspots"
// which can be updated independently and displayed on the device.
type Viewport struct {
	device       *Device
	backingImage *Image
	mode         string
	width        int
	height       int
	x            int
	y            int
	dither       bool
	hotspots     map[hotspotWithXY]any
}

// NewViewport creates a new Viewport with the provided device, width, height,
func NewViewport(device *Device, width int, height int, mode string, dither bool) *Viewport {
	if mode == "" {
		mode = device.mode
	}
	return &Viewport{
		device:       device,
		backingImage: NewImage(mode, width, height, 0),
		mode:         mode,
		width:        width,
		height:       height,
		dither:       dither,
		hotspots:     make(map[hotspotWithXY]any),
	}
}

// Width returns the viewport width in pixels.
func (v *Viewport) Width() int {
	return v.width
}

// Height returns the viewport height in pixels.
func (v *Viewport) Height() int {
	return v.height
}

// Display renders the provided image on the viewport.
func (v *Viewport) Display(image *Image) error {
	if (image.Width() != v.width) || (image.Height() != v.height) {
		return ErrImageSizeMismatch
	}
	v.backingImage.Paste(image)
	v.refresh(true)
	return nil
}

// SetPosition sets the viewport position to the provided coordinates.
func (v *Viewport) SetPosition(x int, y int) {
	v.x = x
	v.y = y
	v.refresh(true)
}

// AddHotspot adds a hotspot to the viewport at the provided coordinates.
func (v *Viewport) AddHotspot(hotspot hotspotable, x int, y int) error {
	if (x < 0) || (y < 0) || (v.Width()-hotspot.Width() < x) || (v.Height()-hotspot.Height() < y) {
		return ErrInvalidHotspot
	}
	v.hotspots[hotspotWithXY{hotspot, x, y}] = struct{}{}
	return nil
}

// RemoveHotspot removes the hotspot from the viewport at the provided
// coordinates and erases the area.
func (v *Viewport) RemoveHotspot(hotspot hotspotable, x int, y int) {
	delete(v.hotspots, hotspotWithXY{hotspot, x, y})
	eraser := NewImage(v.mode, hotspot.Width(), hotspot.Height(), 0)
	defer eraser.Free()
	v.backingImage.PasteXY(eraser, x, y)
}

// ClearHotspots removes all hotspots from the viewport.
func (v *Viewport) ClearHotspots() {
	for hotspotXY := range v.hotspots {
		v.RemoveHotspot(hotspotXY.hotspot, hotspotXY.x, hotspotXY.y)
	}
}

// IsOverlappingViewport returns true if the hotspot at the provided coordinates
// is overlapping with the viewport.
func (v *Viewport) IsOverlappingViewport(hotspot hotspotable, x int, y int) bool {
	l1, t1, r1, b1 := calcBounds(x, y, hotspot)
	l2, t2, r2, b2 := calcBounds(v.x, v.y, v.device)
	return rangeOverlap(l1, r1, l2, r2) && rangeOverlap(t1, b1, t2, b2)
}

// Refresh updates the viewport with the latest hotspots.
func (v *Viewport) Refresh() {
	v.refresh(false)
}

func (v *Viewport) refresh(force bool) {
	var wg sync.WaitGroup
	for hotspotXY := range v.hotspots {
		if hotspotXY.hotspot.ShouldRedraw() && v.IsOverlappingViewport(hotspotXY.hotspot, hotspotXY.x, hotspotXY.y) {
			force = true
			wg.Add(1)
			go func(hotspot hotspotable, x, y int) {
				defer wg.Done()
				hotspot.PasteInto(v.backingImage, x, y)
			}(hotspotXY.hotspot, hotspotXY.x, hotspotXY.y)
		}
	}
	wg.Wait()
	if !force {
		return
	}
	cropped := v.backingImage.Crop(v.cropBox())
	defer cropped.Free()
	if v.dither {
		converted := cropped.Convert(v.device.mode)
		defer converted.Free()
		v.device.Display(converted)
	} else {
		v.device.Display(v.backingImage)
	}
}

func (v *Viewport) cropBox() (int, int, int, int) {
	left, top := v.x, v.y
	right, bottom := left+v.device.Width(), top+v.device.Height()
	return left, top, right, bottom
}

func calcBounds(x int, y int, entity widthHeight) (int, int, int, int) {
	left, top := x, y
	right, bottom := left+entity.Width(), top+entity.Height()
	return left, top, right, bottom
}

func rangeOverlap(aMin, aMax, bMin, bMax int) bool {
	return (aMin < bMax) && (bMin < aMax)
}

// Free releases the resources associated with the viewport.
func (v *Viewport) Free() {
	v.backingImage.Free()
}

// Hotspot represents an area on the display that can be updated independently
// and displayed on the device.
type Hotspot struct {
	width  int
	height int
	drawFn func(*Draw, int, int)
}

// NewHotspot creates a new Hotspot with the provided width, height, and
// function that updates it.
func NewHotspot(width, height int, drawFn func(*Draw, int, int)) *Hotspot {
	return &Hotspot{
		width:  width,
		height: height,
		drawFn: drawFn,
	}
}

// Width returns the hotspot width in pixels.
func (h *Hotspot) Width() int {
	return h.width
}

// Height returns the hotspot height in pixels.
func (h *Hotspot) Height() int {
	return h.height
}

// ShouldRedraw returns true if the hotspot should be redrawn.
func (h *Hotspot) ShouldRedraw() bool {
	return true
}

// PasteInto renders the hotspot on the provided image at the provided
// coordinates.
func (h *Hotspot) PasteInto(image *Image, x, y int) {
	img := NewImage(image.Mode(), h.width, h.height, 0)
	defer img.Free()
	draw := NewDraw(img)
	defer draw.Free()
	h.update(draw)
	image.PasteXY(img, x, y)
}

func (h *Hotspot) update(draw *Draw) {
	h.drawFn(draw, h.width, h.height)
}

// Snapshot represents a snapshot of a hotspot that has a refresh interval.
type Snapshot struct {
	Hotspot
	interval    time.Duration
	lastUpdated time.Time
}

// NewSnapshot creates a new hotspot and wraps it into a snapshot with the
// provided refresh interval.
func NewSnapshot(width, height int, drawFn func(*Draw, int, int), interval time.Duration) *Snapshot {
	return &Snapshot{
		Hotspot:     *NewHotspot(width, height, drawFn),
		interval:    interval,
		lastUpdated: time.Time{},
	}
}

// ShouldRedraw returns true if the snapshot should be redrawn, i.e. more than
// the required interval has passed since the last update.
func (s *Snapshot) ShouldRedraw() bool {
	return time.Since(s.lastUpdated) > s.interval
}

// PasteInto renders the snapshot on the provided image at the provided
// coordinates and updates the last updated time.
func (s *Snapshot) PasteInto(image *Image, x, y int) {
	s.Hotspot.PasteInto(image, x, y)
	s.lastUpdated = time.Now()
}
