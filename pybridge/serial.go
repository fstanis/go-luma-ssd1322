package pybridge

// #include "wrappers.h"
import "C"

type serialModule C.PyObject

func newSerialModule() *serialModule {
	return (*serialModule)(C.import_module(C.SERIAL_MODULE))
}

func (m *serialModule) SPI(port int, device int) *SerialInterface {
	return (*SerialInterface)(C.m_serial_spi((*C.PyObject)(m), C.int(port), C.int(device)))
}

func (m *serialModule) free() {
	C.decref((*C.PyObject)(m))
}

type SerialInterface C.PyObject

func (m *SerialInterface) Free() {
	C.decref((*C.PyObject)(m))
}
