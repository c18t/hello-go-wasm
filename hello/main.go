package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// 標準出力は、デベロッパーツールのConsoleへ出力される
	fmt.Println("Hello, WebAssembly!")

	// GoからJSを操作するにはsyscall/jsのGlobal()を呼ぶ
	window := js.Global()
	div := window.Get("document").Call("getElementById", "hello")
	div.Set("innerText", "Hello, WebAssembly!")
}
