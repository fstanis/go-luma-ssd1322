package pybridge

// #include "wrappers.h"
import "C"
import "unsafe"

type imageModule C.PyObject

func newImageModule() *imageModule {
	return (*imageModule)(C.import_module(C.IMAGE_MODULE))
}

func (m *imageModule) free() {
	C.decref((*C.PyObject)(m))
}

type Image C.PyObject

func (m *imageModule) NewImage(mode string, width int, height int, color int) *Image {
	modeStr := C.CString(mode)
	defer C.free(unsafe.Pointer(modeStr))
	return (*Image)(C.m_image_new_image((*C.PyObject)(m), modeStr, C.int(width), C.int(height), C.int(color)))
}

func (i *Image) Paste(other *Image) {
	C.image_paste((*C.PyObject)(i), (*C.PyObject)(other))
}

func (i *Image) PasteXY(other *Image, x int, y int) {
	C.image_paste_2tuple((*C.PyObject)(i), (*C.PyObject)(other), C.int(x), C.int(y))
}

func (i *Image) Crop(left int, top int, right int, bottom int) *Image {
	return (*Image)(C.image_crop((*C.PyObject)(i), C.int(left), C.int(top), C.int(right), C.int(bottom)))
}

func (i *Image) Convert(mode string) *Image {
	modeStr := C.CString(mode)
	defer C.free(unsafe.Pointer(modeStr))
	return (*Image)(C.image_convert((*C.PyObject)(i), modeStr))
}

func (i *Image) Free() {
	C.decref((*C.PyObject)(i))
}
