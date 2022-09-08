package main

import "fmt"

// 変数宣言 -> 最後に型をつけることを推奨
var x int = 0
var i, j = 1, 2

// 文字列は基本的にダブルクォーテーションで囲む
var y string = "boolはデフォルト"
var y2 = "falseだよん: "

// 最後に型を書くことでリストを宣言できる
var c, python, java bool

func main() {
	python = true
	fmt.Println(x, i, j)
	fmt.Println(y, y2, c, python, java)
}
