package core

import (
	"fmt"
	"math/rand"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (direction Direction) String() string {
	switch direction {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}

type BoardInterface interface {
	MoveBy(dir Direction) bool
	IsSolved() bool
	Shuffle(moves int)
	Snapshot() [][]int
}

type Board struct {
	Size   int
	Tiles  [][]int // 0은 빈 칸
	emptyX int     // 빈 칸의 x좌표
	emptyY int     // 빈 칸의 y좌표
}

func (b *Board) MoveBy(dir Direction) bool {
	switch dir {
	case Up:
		return b.move(0, 1)
	case Down:
		return b.move(0, -1)
	case Left:
		return b.move(1, 0)
	case Right:
		return b.move(-1, 0)
	default:
		return false
	}
}

func (b *Board) move(dx, dy int) bool {
	x, y := b.emptyX+dx, b.emptyY+dy
	if x < 0 || x >= b.Size || y < 0 || y >= b.Size {
		return false
	}
	b.Tiles[b.emptyY][b.emptyX], b.Tiles[y][x] = b.Tiles[y][x], b.Tiles[b.emptyY][b.emptyX]
	b.emptyX, b.emptyY = x, y
	return true
}

func (b *Board) IsSolved() bool {
	val := 1
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if i == b.Size-1 && j == b.Size-1 {
				return b.Tiles[i][j] == 0
			}
			if b.Tiles[i][j] != val {
				return false
			}
			val++
		}
	}
	return true
}

func (b *Board) Shuffle(moves int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	dirs := []Direction{Up, Down, Left, Right}
	for i := 0; i < moves; i++ {
		d := dirs[r.Intn(4)]
		b.MoveBy(d)
	}
}

func (b *Board) Snapshot() [][]int {
	cp := make([][]int, b.Size)
	for i := range b.Tiles {
		cp[i] = make([]int, b.Size)
		copy(cp[i], b.Tiles[i])
	}
	return cp
}

func (b *Board) Print() {
	for _, row := range b.Tiles {
		for _, val := range row {
			if val == 0 {
				fmt.Print("   ")
			} else {
				fmt.Printf("%2d ", val)
			}
		}
		fmt.Println()
	}
}

func NewBoard(size int) *Board {
	b := &Board{
		Size:  size,
		Tiles: make([][]int, size),
	}
	val := 1
	for i := 0; i < size; i++ {
		b.Tiles[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if val == size*size {
				b.Tiles[i][j] = 0 // 빈 칸
				b.emptyX = j
				b.emptyY = i
			} else {
				b.Tiles[i][j] = val
				val++
			}
		}
	}
	return b
}
