//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/meteormin/slide-puzzle/internal/core"
	"github.com/meteormin/slide-puzzle/internal/wasm"
)

func main() {
	ch := make(chan struct{})
	container, err := core.New(2)
	if err != nil {
		return
	}

	controller := wasm.NewController(container)
	controller.Render()

	<-ch // 프로그램 종료 방지
}
