package timer

import (
	"github.com/meteormin/slide-puzzle/internal/core"
	"time"
)

type Timer struct {
	startTime time.Time
	endTime   time.Time
	elapsed   time.Duration
}

func (t *Timer) HandleMove(_ [][]int, _ core.Direction) error {
	if t.startTime.IsZero() {
		t.start()
	}
	return nil
}

func (t *Timer) HandleSolved(_ [][]int) error {
	t.end()
	return nil
}

func (t *Timer) HandleShuffle(_ [][]int, _ int) error {
	t.start()
	return nil
}

func (t *Timer) HandleReset(_ [][]int, _ int) error {
	t.reset()
	return nil
}

func (t *Timer) start() {
	t.startTime = time.Now()
}

func (t *Timer) end() {
	t.endTime = time.Now()
	t.elapsed = t.endTime.Sub(t.startTime)
}

func (t *Timer) reset() {
	t.startTime = time.Time{}
	t.endTime = time.Time{}
	t.elapsed = 0
}

func (t *Timer) Elapsed() time.Duration {
	return t.elapsed
}

func New() *Timer {
	return &Timer{
		startTime: time.Time{},
		endTime:   time.Time{},
		elapsed:   0,
	}
}
