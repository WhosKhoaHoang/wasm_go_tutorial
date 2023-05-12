// Credit to this tutorial
// https://www.youtube.com/watch?v=4kBvvk2Bzis
// Written version:
// https://tutorialedge.net/golang/go-webassembly-tutorial/
// NOTE: The tutorial was created several years ago and some parts of the Go
//		 Go code needed to be updated to conform to the current version's
//		 standards (v1.20.3 at the time of this writing)

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// Needed to change this function to conform to what js.FuncOf()
// expects as a value. Used this SO discussion as reference:
// https://stackoverflow.com/questions/55800163/golangs-syscall-js-js-newcallback-is-undefined
func add(this js.Value, args []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", args[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	//sum := js.ValueOf(args[0].Int() + args[1].Int())
	sum := js.ValueOf(int1 + int2)

	// These two lines were for outputting the result after clicking
	// on the button associated with this function.
	js.Global().Set("output", sum)
	fmt.Println(sum)

	// Output result onto the page instead of to the console
	js.Global().Get("document").Call("getElementById", args[2].String()).Set("value", sum)

	return sum
}

func sub(this js.Value, args []js.Value) interface{} {
	diff := js.ValueOf(args[0].Int() - args[1].Int())
	js.Global().Set("output", diff)
	fmt.Println(diff)
	return diff
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add)) // js.NewCallback is no longer supported!
	js.Global().Set("sub", js.FuncOf(sub))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Go WebAssembly initialized")
	registerCallbacks()
	<-c 
}