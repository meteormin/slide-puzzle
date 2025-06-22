package core

type ContainerInterface interface {
	BoardInterface
	Reset(size int) error
	AddListener(listener Listener)
	ErrorHandler(handler func(error))
}

type Listener interface {
	HandleMove(tiles [][]int, direction Direction) error
	HandleSolved(tiles [][]int) error
	HandleShuffle(tiles [][]int, moves int) error
	HandleReset(tiles [][]int, size int) error
}

type Container struct {
	board        BoardInterface
	listeners    []Listener
	errorHandler func(error)
}

func (c *Container) Reset(size int) error {
	if size < 2 {
		return ErrInvalidSize
	}
	c.board = NewBoard(size)
	for _, listener := range c.listeners {
		if err := listener.HandleReset(c.Snapshot(), size); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) Snapshot() [][]int {
	return c.board.Snapshot()
}

func (c *Container) AddListener(listener Listener) {
	c.listeners = append(c.listeners, listener)
}

func (c *Container) ErrorHandler(handler func(error)) {
	c.errorHandler = handler
}

func (c *Container) MoveBy(direction Direction) bool {
	moved := c.board.MoveBy(direction)
	if moved {
		err := c.onMove(c.Snapshot(), direction)
		if err != nil {
			c.errorHandler(err)
		}
	} else {
		c.errorHandler(ErrInvalidDirection)
	}
	return moved
}

func (c *Container) IsSolved() bool {
	solved := c.board.IsSolved()
	if solved {
		err := c.onSolved(c.Snapshot())
		if err != nil {
			c.errorHandler(err)
		}
	}
	return solved
}

func (c *Container) Shuffle(moves int) {
	c.board.Shuffle(moves)
	err := c.onShuffle(c.Snapshot(), moves)
	if err != nil {
		c.errorHandler(err)
	}
}

func (c *Container) onMove(titles [][]int, direction Direction) error {
	for _, listener := range c.listeners {
		if err := listener.HandleMove(titles, direction); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) onSolved(tiles [][]int) error {
	for _, listener := range c.listeners {
		if err := listener.HandleSolved(tiles); err != nil {
			return err
		}
	}
	return nil
}

func (c *Container) onShuffle(tiles [][]int, moves int) error {
	for _, listener := range c.listeners {
		if err := listener.HandleShuffle(tiles, moves); err != nil {
			return err
		}
	}
	return nil
}

func New(size int) (ContainerInterface, error) {
	if size < 2 {
		return nil, ErrInvalidSize
	}
	return &Container{board: NewBoard(size)}, nil
}
