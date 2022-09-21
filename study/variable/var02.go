package main

import "fmt"

func main() {
	// 関数の中でのみ「:=」の代入文を使い、暗黙的な型宣言ができる
	x := 1
	y, z := 2, 3
	str := "文字列だにょん"
	c, python, java := true, false, "文字列だにょ"

	fmt.Println(x, y+z)
	fmt.Println(str)
	fmt.Println(c, python, java)
}
