package wasm

import (
	"github.com/meteormin/slide-puzzle/internal/core"
	"strconv"
	"strings"
	"syscall/js"
)

const (
	BoardID = "puzzle-board"
)

type Controller struct {
	document     js.Value
	boardSection js.Value
	board        *Board
	coreBoard    core.BoardInterface
}

func (c *Controller) handleClickTile(tileElement js.Value) {
	domID := tileElement.Get("id").String()
	splitID := strings.Split(domID, "-")
	y, errY := strconv.Atoi(splitID[1])
	x, errX := strconv.Atoi(splitID[2])
	if errY != nil || errX != nil {
		js.Global().Call("console.error", "Invalid tile ID:", domID)
		return
	}

	direction := c.board.MoveableDirection(x, y)
	if direction == core.NONE {
		js.Global().Call("console.error", "You can't move tile from this board")
		return
	}

	if !c.coreBoard.MoveBy(direction) {
		js.Global().Call("console.error", "You can't move tile from this board")
	}

	c.board = NewBoard(c.coreBoard.Snapshot())
	c.Render()
}

// Render initializes the board and sets up event listeners for each tile.
func (c *Controller) Render() {
	for _, tile := range c.board.Tiles {
		for _, t := range tile {
			tileElement := c.document.Call("createElement", "div")
			tileElement.Set("id", t.DomID)
			tileElement.Get("classList").Call("add", "tile")
			if t.Empty {
				tileElement.Get("classList").Call("add", "empty")
			} else {
				tileElement.Set("textContent", t.DisplayValue)
			}
			c.boardSection.Call("appendChild", tileElement)
		}
	}
}

func NewController(b core.BoardInterface) *Controller {
	document := js.Global().Get("document")
	boardSection := document.Call("getElementById", BoardID)
	titles := b.Snapshot()

	return &Controller{
		document:     document,
		boardSection: boardSection,
		board:        NewBoard(titles),
		coreBoard:    b,
	}
}
