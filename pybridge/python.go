package pybridge

// #include <Python.h>
import "C"
import "sync/atomic"

var initialized = new(atomic.Bool)

var DeviceModule *deviceModule
var FramebufferModule *framebufferModule
var ImageModule *imageModule
var ImageDrawModule *imageDrawModule
var ImageFontModule *imageFontModule
var SerialModule *serialModule
var VirtualModule *virtualModule

func Init() {
	if !initialized.CompareAndSwap(false, true) {
		return
	}
	C.Py_Initialize()
	DeviceModule = newDeviceModule()
	FramebufferModule = newFramebufferModule()
	ImageModule = newImageModule()
	ImageDrawModule = newImageDrawModule()
	ImageFontModule = newImageFontModule()
	SerialModule = newSerialModule()
	VirtualModule = newVirtualModule()
}

func Finalize() {
	if !initialized.Load() {
		return
	}
	DeviceModule.free()
	FramebufferModule.free()
	ImageModule.free()
	ImageDrawModule.free()
	ImageFontModule.free()
	SerialModule.free()
	VirtualModule.free()
	C.Py_Finalize()
}
