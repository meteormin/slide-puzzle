package counter

import "github.com/meteormin/slide-puzzle/internal/core"

type Counter struct {
	count int
}

func (c *Counter) HandleMove(_ [][]int, _ core.Direction) error {
	c.count++
	return nil
}

func (c *Counter) HandleSolved(_ [][]int) error {
	// Do nothing, just a counter
	return nil
}

func (c *Counter) HandleShuffle(_ [][]int, _ int) error {
	c.count = 0 // Reset the counter on shuffle
	return nil
}

func (c *Counter) HandleReset(_ [][]int, _ int) error {
	c.count = 0 // Reset the counter on reset
	return nil
}

func (c *Counter) Count() int {
	return c.count
}

func New() *Counter {
	return &Counter{count: 0}
}
