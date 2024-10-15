package ssd1322

import (
	"errors"

	"github.com/fstanis/go-luma-ssd1322/pybridge"
)

var (
	ErrFailedToLoadSerial   = errors.New("failed to load serial")
	ErrFailedToLoadDevice   = errors.New("failed to load device")
	ErrFailedToLoadFont     = errors.New("failed to load font")
	ErrFailedToLoadTerminal = errors.New("failed to load terminal")
	ErrInvalidFont          = errors.New("invalid font")
)

// Device is a wrapper around the SSD1322 device.
type Device struct {
	device *pybridge.Device
	fonts  map[string]*Font
	width  int
	height int
	mode   string
}

// NewDevice attempts to establish a connection to the provided serial port with
// the given width and height (in pixels) as awell as a mode, such as "RGB".
func NewDevice(port int, width int, height int, mode string) (*Device, error) {
	serial := pybridge.SerialModule.SPI(port, 0)
	if serial == nil {
		return nil, ErrFailedToLoadSerial
	}
	defer serial.Free()
	if width == 0 {
		width = 256
	}
	if height == 0 {
		height = 64
	}
	if mode == "" {
		mode = "RBG"
	}
	device := pybridge.DeviceModule.SSD1322(serial, width, height, 2, mode, nil)
	if device == nil {
		return nil, ErrFailedToLoadDevice
	}
	return &Device{
		device: device,
		fonts:  make(map[string]*Font),
		width:  width,
		height: height,
		mode:   mode,
	}, nil
}

// Width returns the device width in pixels.
func (d *Device) Width() int {
	return d.width
}

// Height returns the device height in pixels.
func (d *Device) Height() int {
	return d.height
}

// Display displays the provided image on the device.
func (d *Device) Display(image *Image) {
	d.device.Display(image.img)
}

// Font represents an image font loaded on the device.
type Font struct {
	font   *pybridge.ImageFont
	Family string
	Style  string
	Path   string
}

// GetBBox returns the bounding box of the provided text.
func (f *Font) GetBBox(text string) (int, int, int, int) {
	return f.font.GetBBox(text)
}

// LoadFont attempts to load a font from the provided path with the given font
// size.
func (d *Device) LoadFont(path string, size int) (*Font, error) {
	if current, exists := d.fonts[path]; exists {
		return current, nil
	}
	font := pybridge.ImageFontModule.Truetype(path, size, 0, "", "BASIC")
	if font == nil {
		return nil, ErrFailedToLoadFont
	}
	name := font.GetName()
	newfont := &Font{
		font:   font,
		Family: name.Family,
		Style:  name.Style,
		Path:   path,
	}
	d.fonts[path] = newfont
	return newfont, nil
}

// Free frees a previously loaded font.
func (d *Device) Free() {
	for _, font := range d.fonts {
		font.font.Free()
	}
	d.device.Free()
}

// Terminal creates an instance of a terminal object with the provided font.
func (d *Device) Terminal(font *Font) (*Terminal, error) {
	term := pybridge.VirtualModule.NewTerminal(d.device, font.font, 255, 0, 4, 0, true, false)
	if term == nil {
		return nil, ErrFailedToLoadTerminal
	}
	return &Terminal{
		term: term,
	}, nil
}
