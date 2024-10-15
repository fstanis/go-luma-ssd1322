#ifndef WRAPPERS_H_
#define WRAPPERS_H_

#include <Python.h>

extern const char* SERIAL_MODULE;
extern const char* DEVICE_MODULE;
extern const char* VIRTUAL_MODULE;
extern const char* MIXIN_MODULE;
extern const char* FRAMEBUFFER_MODULE;
extern const char* IMAGEFONT_MODULE;
extern const char* IMAGEDRAW_MODULE;
extern const char* IMAGE_MODULE;

extern const char* LAYOUT;
extern const char* BASIC;
extern const char* RAQM;

extern const char* EMPTY;

void decref(PyObject* obj);

PyObject* import_module(const char* name);

PyObject* m_serial_spi(PyObject* module, int port, int device);

PyObject* m_device_ssd1322(PyObject* module,
                           PyObject* serial,
                           int width,
                           int height,
                           int rotate,
                           const char* mode,
                           PyObject* framebuffer);

PyObject* device_display(PyObject* device, PyObject* image);

PyObject* m_imagefont_truetype(PyObject* module,
                               const char* font_path,
                               int font_size,
                               int index,
                               const char* encoding,
                               PyObject* layout_engine);
PyObject* imagefont_getname(PyObject* font);
PyObject* imagefont_getbbox(PyObject* font, const char* text);
PyObject* m_virtual_terminal(PyObject* module,
                             PyObject* device,
                             PyObject* font,
                             int color,
                             int bgcolor,
                             int tabstop,
                             int line_height,
                             PyObject* animate,
                             PyObject* word_wrap);

void terminal_clear(PyObject* terminal);

void terminal_println(PyObject* terminal, const char* text);

void terminal_puts(PyObject* terminal, const char* text);

void terminal_putch(PyObject* terminal, int ch);

PyObject* m_mixin_capabilities(PyObject* module,
                               int width,
                               int height,
                               int rotate,
                               const char* mode);

PyObject* m_image_new_image(PyObject* module,
                            const char* mode,
                            int width,
                            int height,
                            int color);

void image_paste(PyObject* image, PyObject* target);

void image_paste_2tuple(PyObject* image, PyObject* target, int x, int y);

PyObject* image_crop(PyObject* image, int left, int top, int right, int bottom);

PyObject* image_convert(PyObject* image, const char* mode);

PyObject* m_imagedraw_draw(PyObject* module, PyObject* image);

void imagedraw_text(PyObject* imagedraw,
                    int x,
                    int y,
                    const char* text,
                    int fill,
                    PyObject* font);

void imagedraw_bitmap(PyObject* imagedraw,
                      int x,
                      int y,
                      PyObject* bitmap,
                      int fill);
void imagedraw_rectangle(PyObject* imagedraw,
                         int x1,
                         int y1,
                         int x2,
                         int y2,
                         int fill);
PyObject* m_framebuffer_diff_to_previous(PyObject* module, int num_segments);

PyObject* m_framebuffer_full_frame(PyObject* module);

#endif  // WRAPPERS_H_
