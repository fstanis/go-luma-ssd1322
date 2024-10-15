package ssd1322

import "github.com/fstanis/go-luma-ssd1322/pybridge"

// Terminal represents an object that can be used to output text to the display.
type Terminal struct {
	term *pybridge.Terminal
}

// Clear clears all text on the terminal.
func (t *Terminal) Clear() {
	t.term.Clear()
}

// Print prints the provided text with a new line to the terminal.
func (t *Terminal) Println(text string) {
	t.term.Println(text)
}

// Putc prints the provided text to the terminal.
func (t *Terminal) Puts(text string) {
	t.term.Puts(text)
}

// Putch prints the provided character to the terminal.
func (t *Terminal) Putch(ch rune) {
	t.term.Putch(ch)
}

// Free releases the resources associated with the terminal.
func (t *Terminal) Free() {
	t.term.Free()
}
