package main

import "fmt"

func main() {
	sum := 1
	// 初期化と後処理ステートメントの記述は任意.「;」も省略可 -> while文は下記のように記述する.
	// ループ条件も省略すれば無限ループになる
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
