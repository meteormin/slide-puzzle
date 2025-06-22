package logger

import (
	"github.com/meteormin/slide-puzzle/internal/core"
	"go.uber.org/zap"
)

type Listener struct {
	logger *zap.SugaredLogger
}

func (l *Listener) HandleReset(_ [][]int, size int) error {
	l.logger.Infof("Reset Puzzle of size %d", size)
	return nil
}

func (l *Listener) HandleMove(_ [][]int, direction core.Direction) error {
	l.logger.Infof("Move %s", direction)
	return nil
}

func (l *Listener) HandleSolved(_ [][]int) error {
	l.logger.Infof("Solved Puzzle")
	return nil
}

func (l *Listener) HandleShuffle(_ [][]int, moves int) error {
	l.logger.Infof("Shuffled Puzzle with %d moves", moves)
	return nil
}

func NewListener(logger *zap.SugaredLogger) *Listener {
	return &Listener{logger: logger}
}
