package main

import (
	"fmt"
	"syscall/js"
)

// go 1.25からはエクスポートが書けるとのこと
//
//go:wasmexport fib
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	// 今はJS側に生やすコードが必要
	js.Global().Set("fib", js.FuncOf(func(this js.Value, args []js.Value) any {
		return fib(args[0].Int())
	}))

	num := 10
	fmt.Printf("fib(%d) is %d\n", num, fib(num))
	select {}
}
