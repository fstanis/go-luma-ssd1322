package pybridge

// #include "wrappers.h"
import "C"
import (
	"unsafe"
)

type deviceModule C.PyObject

func newDeviceModule() *deviceModule {
	return (*deviceModule)(C.import_module(C.DEVICE_MODULE))
}

func (m *deviceModule) SSD1322(serial *SerialInterface, width int, height int, rotate int, mode string, framebuffer *Framebuffer) *Device {
	if width == 0 || height == 0 || mode == "" {
		return nil
	}
	if framebuffer == nil {
		framebuffer = (*Framebuffer)(C.Py_None)
	}
	modeStr := C.CString(mode)
	defer C.free(unsafe.Pointer(modeStr))
	return (*Device)(C.m_device_ssd1322((*C.PyObject)(m), (*C.PyObject)(serial), C.int(width), C.int(height), C.int(rotate), modeStr, (*C.PyObject)(framebuffer)))
}

func (m *deviceModule) free() {
	C.decref((*C.PyObject)(m))
}

type Device C.PyObject

func (m *Device) Display(image *Image) {
	C.device_display((*C.PyObject)(m), (*C.PyObject)(image))
}

func (m *Device) Free() {
	C.decref((*C.PyObject)(m))
}
