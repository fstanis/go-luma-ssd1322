package pybridge

// #include "wrappers.h"
import "C"
import "unsafe"

type imageDrawModule C.PyObject

func newImageDrawModule() *imageDrawModule {
	return (*imageDrawModule)(C.import_module(C.IMAGEDRAW_MODULE))
}

func (m *imageDrawModule) free() {
	C.decref((*C.PyObject)(m))
}

type ImageDraw C.PyObject

func (m *imageDrawModule) Draw(image *Image) *ImageDraw {
	return (*ImageDraw)(C.m_imagedraw_draw((*C.PyObject)(m), (*C.PyObject)(image)))
}

func (i *ImageDraw) Text(x int, y int, text string, fill int, font *ImageFont) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.imagedraw_text((*C.PyObject)(i), C.int(x), C.int(y), textStr, C.int(fill), (*C.PyObject)(font))
}

func (i *ImageDraw) Bitmap(x int, y int, bitmap *Image, fill int) {
	C.imagedraw_bitmap((*C.PyObject)(i), C.int(x), C.int(y), (*C.PyObject)(bitmap), C.int(fill))
}

func (i *ImageDraw) Rectangle(x1 int, y1 int, x2 int, y2 int, fill int) {
	C.imagedraw_rectangle((*C.PyObject)(i), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(fill))
}

func (i *ImageDraw) Free() {
	C.decref((*C.PyObject)(i))
}
