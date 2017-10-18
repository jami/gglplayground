package system

import "time"

// Timer helper
type Timer struct {
	currentTime int64
	lastTime    int64
	tick        float64
}

// Reset sync the timer
func (t *Timer) Reset() {
	t.currentTime = t.getCurrentTimeMs()
	t.lastTime = t.currentTime
	t.tick = 0
}

// Update updates the timer
func (t *Timer) Update() {
	t.currentTime = t.getCurrentTimeMs()
	delta := t.currentTime - t.lastTime

	t.tick = float64(delta) * 0.001
	t.lastTime = t.currentTime
}

// GetTick returns the duration in s
func (t *Timer) GetTick() float64 {
	return t.tick
}

func (t *Timer) getCurrentTimeMs() int64 {
	return time.Now().UnixNano() / 1000000
}

// NewTimer creates a new timer
func NewTimer() *Timer {
	return &Timer{}
}
