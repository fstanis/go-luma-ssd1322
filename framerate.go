package ssd1322

import (
	"math"
	"time"
)

// FrameRateRegulator is used to throttle calls based on the provided framerate.
type FramerateRegulator struct {
	maxSleep         time.Duration
	enterTime        time.Time
	startTime        time.Time
	lastTime         time.Time
	totalTransitTime time.Duration
	called           int
}

// NewFramerateRegulator creates a new FramerateRegulator with the provided
// target framerate.
func NewFramerateRegulator(fps float64) *FramerateRegulator {
	return &FramerateRegulator{
		maxSleep:  time.Duration(math.Round(float64(time.Second.Nanoseconds()) / fps)),
		enterTime: time.Time{},
	}
}

// Enter is called before drawing on the screen to signify the start of a frame.
func (f *FramerateRegulator) Enter() {
	f.enterTime = time.Now()
	if f.startTime.IsZero() {
		f.startTime = f.enterTime
		f.lastTime = f.enterTime
	}
}

// Exit is called after drawing on the screen to signify the end of a frame. It
// may result in a sleep to throttle the framerate.
func (f *FramerateRegulator) Exit() {
	f.called++
	f.totalTransitTime = time.Since(f.enterTime)
	if f.maxSleep >= 0 {
		elapsed := time.Since(f.lastTime)
		sleepFor := f.maxSleep - elapsed
		if sleepFor > 0 {
			time.Sleep(sleepFor)
		}
	}
	f.lastTime = time.Now()
}

// EffectiveFPS returns the effective FPS based on previous calls to Enter and
// Exit.
func (f *FramerateRegulator) EffectiveFPS() float64 {
	elapsed := time.Since(f.startTime)
	return float64(f.called) / elapsed.Seconds()
}

// AverageTransitTime returns the average time taken to process a frame based on
// previous calls to Enter and Exit.
func (f *FramerateRegulator) AverageTransitTime() time.Duration {
	if f.called == 0 {
		return 0
	}
	return f.totalTransitTime / time.Duration(f.called)
}
