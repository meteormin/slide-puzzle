package wasm

import (
	"github.com/meteormin/slide-puzzle/internal/core"
	"syscall/js"
)

type Controller struct {
	board    core.BoardInterface
	document js.Value
}

func (c *Controller) MoveBy(dir core.Direction) {

}

func NewController(board core.BoardInterface) *Controller {
	return &Controller{
		board:    board,
		document: js.Global().Get("document"),
	}
}
