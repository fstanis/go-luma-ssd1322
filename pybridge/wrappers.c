#include "wrappers.h"

const char* SERIAL_MODULE = "luma.core.interface.serial";
const char* DEVICE_MODULE = "luma.oled.device";
const char* VIRTUAL_MODULE = "luma.core.virtual";
const char* MIXIN_MODULE = "luma.core.mixin";
const char* FRAMEBUFFER_MODULE = "luma.core.framebuffer";
const char* IMAGEFONT_MODULE = "PIL.ImageFont";
const char* IMAGEDRAW_MODULE = "PIL.ImageDraw";
const char* IMAGE_MODULE = "PIL.Image";

const char* LAYOUT = "Layout";
const char* BASIC = "BASIC";
const char* RAQM = "RAQM";

const char* EMPTY = "";

void decref(PyObject* obj) {
  Py_DECREF(obj);
}

PyObject* import_module(const char* name) {
  PyObject* py_name = PyUnicode_FromString(name);
  PyObject* module = PyImport_Import(py_name);
  Py_DECREF(py_name);
  return module;
}

PyObject* m_serial_spi(PyObject* module, int port, int device) {
  return PyObject_CallMethod(module, "spi", "(OOii)", Py_None, Py_None, port,
                             device);
}

PyObject* m_device_ssd1322(PyObject* module,
                           PyObject* serial,
                           int width,
                           int height,
                           int rotate,
                           const char* mode,
                           PyObject* framebuffer) {
  return PyObject_CallMethod(module, "ssd1322", "(OiiisO)", serial, width,
                             height, rotate, mode, Py_None);
}

PyObject* device_display(PyObject* device, PyObject* image) {
  return PyObject_CallMethod(device, "display", "(O)", image);
}

PyObject* m_imagefont_truetype(PyObject* module,
                               const char* font_path,
                               int font_size,
                               int index,
                               const char* encoding,
                               PyObject* layout_engine) {
  return PyObject_CallMethod(module, "truetype", "(siisO)", font_path,
                             font_size, index, encoding, layout_engine);
}

PyObject* imagefont_getname(PyObject* font) {
  return PyObject_CallMethod(font, "getname", NULL);
}

PyObject* imagefont_getbbox(PyObject* font, const char* text) {
  return PyObject_CallMethod(font, "getbbox", "(s)", text);
}

PyObject* m_virtual_terminal(PyObject* module,
                             PyObject* device,
                             PyObject* font,
                             int color,
                             int bgcolor,
                             int tabstop,
                             int line_height,
                             PyObject* animate,
                             PyObject* word_wrap) {
  return PyObject_CallMethod(module, "terminal", "(OOiiiiOO)", device, font,
                             color, bgcolor, tabstop, line_height, animate,
                             word_wrap);
}

void terminal_clear(PyObject* terminal) {
  PyObject_CallMethod(terminal, "clear", NULL);
}

void terminal_println(PyObject* terminal, const char* text) {
  PyObject_CallMethod(terminal, "println", "(s)", text);
}

void terminal_puts(PyObject* terminal, const char* text) {
  PyObject_CallMethod(terminal, "puts", "(s)", text);
}

void terminal_putch(PyObject* terminal, int ch) {
  PyObject_CallMethod(terminal, "putch", "(C)", ch);
}

PyObject* m_mixin_capabilities(PyObject* module,
                               int width,
                               int height,
                               int rotate,
                               const char* mode) {
  return PyObject_CallMethod(module, "capabilities", "(iiis)", width, height,
                             rotate, mode);
}

PyObject* m_image_new_image(PyObject* module,
                            const char* mode,
                            int width,
                            int height,
                            int color) {
  return PyObject_CallMethod(module, "new", "(s(ii)i)", mode, width, height,
                             color);
}

void image_paste(PyObject* image, PyObject* target) {
  PyObject_CallMethod(image, "paste", "(O)", target);
}

void image_paste_2tuple(PyObject* image, PyObject* target, int x, int y) {
  PyObject_CallMethod(image, "paste", "(O(ii))", target, x, y);
}

PyObject* image_crop(PyObject* image,
                     int left,
                     int top,
                     int right,
                     int bottom) {
  return PyObject_CallMethod(image, "crop", "((iiii))", left, top, right,
                             bottom);
}

PyObject* image_convert(PyObject* image, const char* mode) {
  return PyObject_CallMethod(image, "convert", "(s)", mode);
}

PyObject* m_imagedraw_draw(PyObject* module, PyObject* image) {
  return PyObject_CallMethod(module, "Draw", "(O)", image);
}

void imagedraw_text(PyObject* imagedraw,
                    int x,
                    int y,
                    const char* text,
                    int fill,
                    PyObject* font) {
  PyObject_CallMethod(imagedraw, "text", "((ii)siO)", x, y, text, fill, font);
}

void imagedraw_bitmap(PyObject* imagedraw,
                      int x,
                      int y,
                      PyObject* bitmap,
                      int fill) {
  PyObject_CallMethod(imagedraw, "bitmap", "((ii)Oi)", x, y, bitmap, fill);
}

void imagedraw_rectangle(PyObject* imagedraw,
                         int x1,
                         int y1,
                         int x2,
                         int y2,
                         int fill) {
  PyObject_CallMethod(imagedraw, "rectangle", "((iiii)i)", x1, y1, x2, y2,
                      fill);
}

PyObject* m_framebuffer_diff_to_previous(PyObject* module, int num_segments) {
  return PyObject_CallMethod(module, "diff_to_previous", "(i)", num_segments);
}

PyObject* m_framebuffer_full_frame(PyObject* module) {
  return PyObject_CallMethod(module, "full_frame", NULL);
}
