package main

import "fmt"

func main() {
	z := fruit{"apple", 3}
	// structのフィールドは「.」を用いてアクセスする
	fmt.Println(z.x, z.y)
	// フィールドに値を再代入
	z.x = "orange"
	z.y = 21
	fmt.Println(z.x, z.y)
}
