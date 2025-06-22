package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_IsSolved(t *testing.T) {
	type fields struct {
		Size   int
		Tiles  [][]int
		emptyX int
		emptyY int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "solved board",
			fields: fields{
				Size: 3,
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0},
				},
				emptyX: 2,
				emptyY: 2,
			},
			want: true,
		},
		{
			name: "unsolved board",
			fields: fields{
				Size: 2,
				Tiles: [][]int{
					{1, 2},
					{0, 3},
				},
				emptyX: 0,
				emptyY: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Size:   tt.fields.Size,
				Tiles:  tt.fields.Tiles,
				emptyX: tt.fields.emptyX,
				emptyY: tt.fields.emptyY,
			}
			assert.Equal(t, tt.want, b.IsSolved(),
				"IsSolved() should return %v for board %v", tt.want, b.Tiles)
		})
	}
}

func TestBoard_MoveBy(t *testing.T) {
	type fields struct {
		Size   int
		Tiles  [][]int
		emptyX int
		emptyY int
	}
	type args struct {
		dir Direction
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantTiles [][]int
	}{
		{
			name: "solved, move up",
			fields: fields{
				Size: 3,
				Tiles: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 0},
				},
				emptyX: 2,
				emptyY: 2,
			},
			args: args{
				dir: Up,
			},
			want: false,
			wantTiles: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
		},
		{
			name: "move up",
			fields: fields{
				Size: 2,
				Tiles: [][]int{
					{3, 0},
					{2, 1},
				},
				emptyX: 1,
				emptyY: 0,
			},
			args: args{
				dir: Up,
			},
			want: true,
			wantTiles: [][]int{
				{3, 1},
				{2, 0},
			},
		},
		{
			name: "move down",
			fields: fields{
				Size: 2,
				Tiles: [][]int{
					{3, 1},
					{2, 0},
				},
				emptyX: 1,
				emptyY: 1,
			},
			args: args{
				dir: Down,
			},
			want: true,
			wantTiles: [][]int{
				{3, 0},
				{2, 1},
			},
		},
		{
			name: "move left",
			fields: fields{
				Size: 2,
				Tiles: [][]int{
					{0, 3},
					{2, 1},
				},
				emptyX: 0,
				emptyY: 0,
			},
			args: args{
				dir: Left,
			},
			want: true,
			wantTiles: [][]int{
				{3, 0},
				{2, 1},
			},
		},
		{
			name: "move right",
			fields: fields{
				Size: 2,
				Tiles: [][]int{
					{3, 0},
					{2, 1},
				},
				emptyX: 1,
				emptyY: 0,
			},
			args: args{
				dir: Right,
			},
			want: true,
			wantTiles: [][]int{
				{0, 3},
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Size:   tt.fields.Size,
				Tiles:  tt.fields.Tiles,
				emptyX: tt.fields.emptyX,
				emptyY: tt.fields.emptyY,
			}
			assert.Equal(t, tt.want, b.MoveBy(tt.args.dir),
				"MoveBy should return %v for board %v with direction %v", tt.want, b.Tiles, tt.args.dir)
			assert.Equal(t, tt.wantTiles, b.Tiles, "Tiles should match after move operation")
		})
	}
}

func TestBoard_Shuffle(t *testing.T) {
	type fields struct {
		Size   int
		Tiles  [][]int
		emptyX int
		emptyY int
	}
	type args struct {
		moves int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "shuffle with 10 moves",
			fields: fields{
				Size:   3,
				Tiles:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}},
				emptyX: 2,
				emptyY: 2,
			},
			args: args{
				moves: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				Size:   tt.fields.Size,
				Tiles:  tt.fields.Tiles,
				emptyX: tt.fields.emptyX,
				emptyY: tt.fields.emptyY,
			}
			before := b.Snapshot()
			b.Shuffle(tt.args.moves)
			assert.NotEqual(t, before, b.Tiles, "Shuffle should change the board state after %d moves", tt.args.moves)
		})
	}
}
