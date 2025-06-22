package printer

import (
	"bufio"
	"fmt"
	"github.com/meteormin/slide-puzzle/internal/core"
	"os"
)

type Printer struct {
	w *bufio.Writer
}

func (p *Printer) HandleReset(tiles [][]int, size int) error {
	err := p.Clear()
	_, err = fmt.Fprintf(p.w, "New Puzzle of size %d\n", size)
	if err != nil {
		return err
	}
	return p.WriteTiles(tiles)
}

func (p *Printer) HandleMove(tiles [][]int, _ core.Direction) error {
	err := p.Clear()
	if err != nil {
		return err
	}

	return p.WriteTiles(tiles)
}

func (p *Printer) HandleSolved(tiles [][]int) error {
	err := p.Clear()
	if err != nil {
		return err
	}

	err = p.WriteTiles(tiles)
	if err != nil {
		return err
	}

	_, err = p.w.WriteString("Puzzle Solved!\n")
	return err
}

func (p *Printer) HandleShuffle(tiles [][]int, moves int) error {
	err := p.Clear()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(p.w, "Shuffled Puzzle with %d moves\n", moves)
	if err != nil {
		return err
	}

	return p.WriteTiles(tiles)
}

func (p *Printer) Clear() error {
	_, err := p.w.Write([]byte("\033[2J")) // Clear the screen
	if err != nil {
		return err
	}
	_, err = p.w.Write([]byte("\033[5;1H")) // 커서를 5행 1열로 이동
	if err != nil {
		return err
	}
	return nil
}

func (p *Printer) WriteTiles(tiles [][]int) error {
	for _, row := range tiles {
		for _, val := range row {
			if val == 0 {
				_, err := p.w.WriteString("   ")
				if err != nil {
					return err
				}
			} else {
				_, err := fmt.Fprintf(p.w, "%2d ", val)
				if err != nil {
					return err
				}
			}
		}
		err := p.w.WriteByte(byte('\n'))
		if err != nil {
			return err
		}
	}

	_, err := p.w.WriteString("\n")
	if err != nil {
		return err
	}

	return p.w.Flush()
}

func (p *Printer) Close() error {
	return p.w.Flush()
}

func NewPrinter() *Printer {
	return &Printer{
		w: bufio.NewWriter(os.Stdout),
	}
}
