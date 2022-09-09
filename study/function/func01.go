package main

import "fmt"

// 　引き数が同じ型の場合は最後の型を残して省略できる
func add(x, y, z int) int {
	return (x + y) * z
}

func main() {
	fmt.Println(add(42, 13, 2))
}
