# go-luma-ssd1322

This package is a Go library wrapper around [luma.oled](https://github.com/rm-hull/luma.oled),
specifically the parts used to interface with an SSD1322 OLED display.

It can be used to port something like [train-departure-display](https://github.com/chrisys/train-departure-display)
to Go.

This library uses cgo to interface with luma.oled's Python code, which is
not the most efficient way of interfacing with this screen. For more serious
projects, consider using one that uses [Periph](https://periph.io/) or a similar
project that's native Go.

## Building

### Prerequisites

#### Debian-based

```bash
$ apt-get install python3-dev python3-pip libfreetype6-dev libjpeg-dev build-essential
$ pip install --upgrade luma.oled
```

#### Alpine

```bash
$ apk add python3-dev py3-pip freetype-dev libjpeg-turbo-dev
$ pip install --upgrade luma.oled
```

### CGO flags

Valid `CFLAGS` and `LDFLAGS` must be used for embedding Python.

If `pkg-config` is installed, you can set them as environmental variables:

```bash
export CGO_CFLAGS="$(pkg-config --cflags python3-embed)"
export CGO_LDFLAGS="$(pkg-config --libs python3-embed)"
```

## Example usage

```go
import (
	"github.com/fstanis/go-luma-ssd1322"
	"github.com/fstanis/go-luma-ssd1322/pybridge"
)

func helloWorld() error {
	pybridge.Init()
	defer pybridge.Finalize()

	device, err := ssd1322.NewDevice(0, 0, 0, "1")
	if err != nil {
		return err
	}
	defer device.Free()

	terminal, err = device.Terminal(selectedFont)
	if err != nil {
		return err
	}
	defer terminal.Free()

	terminal.Println("Hello, world!")
}
```
