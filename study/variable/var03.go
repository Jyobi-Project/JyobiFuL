package main

import "fmt"

func main() {
	// デフォルト -> 0
	var i int
	var f float64
	// デフォルト -> false
	var b bool
	// デフォルト -> ""
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
