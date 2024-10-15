package pybridge

// #include "wrappers.h"
import "C"
import (
	"unsafe"
)

type virtualModule C.PyObject

func newVirtualModule() *virtualModule {
	return (*virtualModule)(C.import_module(C.VIRTUAL_MODULE))
}

func (m *virtualModule) NewTerminal(device *Device, font *ImageFont, color int, bgcolor int, tabstop int, lineHeight int, animate bool, wordWrap bool) *Terminal {
	if tabstop == 0 {
		tabstop = 4
	}
	return (*Terminal)(C.m_virtual_terminal((*C.PyObject)(m), (*C.PyObject)(device), (*C.PyObject)(font), C.int(color), C.int(bgcolor), C.int(tabstop), C.int(lineHeight), boolToPyBool(animate), boolToPyBool(wordWrap)))
}

func (m *virtualModule) free() {
	C.decref((*C.PyObject)(m))
}

type Terminal C.PyObject

func (t *Terminal) Clear() {
	C.terminal_clear((*C.PyObject)(t))
}

func (t *Terminal) Println(text string) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.terminal_println((*C.PyObject)(t), textStr)
}

func (t *Terminal) Puts(text string) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.terminal_puts((*C.PyObject)(t), textStr)
}

func (t *Terminal) Putch(ch rune) {
	C.terminal_putch((*C.PyObject)(t), C.int(int(ch)))
}

func (m *Terminal) Free() {
	C.decref((*C.PyObject)(m))
}

func boolToPyBool(b bool) *C.PyObject {
	if b {
		return C.Py_True
	}
	return C.Py_False
}
