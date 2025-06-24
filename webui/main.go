//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func alertCallback(this js.Value, p []js.Value) interface{} {
	js.Global().Call("alert", "버튼이 클릭되었습니다!")
	return nil
}

func main() {
	c := make(chan struct{})

	// 버튼에 이벤트 핸들러 등록
	document := js.Global().Get("document")
	button := document.Call("getElementById", "myButton")
	button.Call("addEventListener", "click", js.FuncOf(alertCallback))

	<-c // 프로그램 종료 방지
}
