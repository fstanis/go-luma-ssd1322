package pybridge

// #include "wrappers.h"
import "C"
import (
	"unsafe"
)

type imageFontModule struct {
	module     *C.PyObject
	layoutEnum map[string]*C.PyObject
}

func newImageFontModule() *imageFontModule {
	module := C.import_module(C.IMAGEFONT_MODULE)
	layout := C.PyObject_GetAttrString(module, C.LAYOUT)
	defer C.decref(layout)
	return &imageFontModule{
		module: module,
		layoutEnum: map[string]*C.PyObject{
			"BASIC": C.PyObject_GetAttrString(layout, C.BASIC),
			"RAQM":  C.PyObject_GetAttrString(layout, C.RAQM),
		},
	}
}

func (m *imageFontModule) Truetype(path string, size int, index int, encoding string, layout string) *ImageFont {
	pathStr := C.CString(path)
	defer C.free(unsafe.Pointer(pathStr))
	if size == 0 {
		size = 10
	}
	layoutObj := m.layoutEnum[layout]
	if layoutObj == nil {
		layoutObj = C.Py_None
	}
	encodingStr := C.EMPTY
	if encoding != "" {
		encodingStr = C.CString(encoding)
		defer C.free(unsafe.Pointer(encodingStr))
	}
	return (*ImageFont)(C.m_imagefont_truetype(m.module, pathStr, C.int(size), C.int(index), encodingStr, layoutObj))
}

func (m *imageFontModule) free() {
	C.decref((*C.PyObject)(m.layoutEnum["BASIC"]))
	C.decref((*C.PyObject)(m.layoutEnum["RAQM"]))
	C.decref((*C.PyObject)(m.module))
}

type ImageFont C.PyObject

type FontName struct {
	Family string
	Style  string
}

func (f *ImageFont) GetName() (result FontName) {
	tuple := C.imagefont_getname((*C.PyObject)(f))
	defer C.decref(tuple)
	familyObj := C.PyTuple_GetItem(tuple, 0)
	styleObj := C.PyTuple_GetItem(tuple, 1)
	if familyObj != C.Py_None {
		familyStr := C.PyUnicode_AsUTF8(familyObj)
		result.Family = C.GoString(familyStr)
	}
	if styleObj != C.Py_None {
		styleStr := C.PyUnicode_AsUTF8(styleObj)
		result.Style = C.GoString(styleStr)
	}
	return
}

func (f *ImageFont) GetBBox(text string) (int, int, int, int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	tuple := C.imagefont_getbbox((*C.PyObject)(f), textStr)
	defer C.decref(tuple)
	leftObj := C.PyTuple_GetItem(tuple, 0)
	rightObj := C.PyTuple_GetItem(tuple, 1)
	topObj := C.PyTuple_GetItem(tuple, 2)
	bottomObj := C.PyTuple_GetItem(tuple, 3)
	left := C.PyLong_AsLong(leftObj)
	right := C.PyLong_AsLong(rightObj)
	top := C.PyLong_AsLong(topObj)
	bottom := C.PyLong_AsLong(bottomObj)
	return (int)(left), (int)(right), (int)(top), (int)(bottom)
}

func (f *ImageFont) Free() {
	C.decref((*C.PyObject)(f))
}
