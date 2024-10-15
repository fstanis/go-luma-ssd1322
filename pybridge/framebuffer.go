package pybridge

// #include "wrappers.h"
import "C"

type framebufferModule struct {
	module    *C.PyObject
	fullFrame *Framebuffer
}

func newFramebufferModule() *framebufferModule {
	return &framebufferModule{
		module:    C.import_module(C.FRAMEBUFFER_MODULE),
		fullFrame: nil,
	}
}

func (m *framebufferModule) DiffToPrevious(numSegments int) *Framebuffer {
	if numSegments == 0 {
		numSegments = 4
	}
	return (*Framebuffer)(C.m_framebuffer_diff_to_previous(m.module, C.int(numSegments)))
}

func (m *framebufferModule) FullFrame() *Framebuffer {
	if m.fullFrame == nil {
		m.fullFrame = (*Framebuffer)(C.m_framebuffer_full_frame(m.module))
	}
	return m.fullFrame
}

func (m *framebufferModule) free() {
	if m.fullFrame != nil {
		m.fullFrame.Free()
	}
	C.decref(m.module)
}

type Framebuffer C.PyObject

func (m *Framebuffer) Free() {
	C.decref((*C.PyObject)(m))
}
