package main

import "fmt"

func main() {
	apple, orange := 42, 2700
	// 「&」をつけてappleのポインタを引き出してpに代入する
	p := &apple
	// ポインタpの値を表示
	fmt.Println(p)
	// 「*」をつけてポインタpを通してappleから値を読み出す
	fmt.Println(*p)
	// 「*」をつけてポインタpを通してappleへ値を再代入する
	*p = 21
	// 再代入した値を読み出す
	fmt.Println(*p)

	// orangeのポインタを引き出してポインタpに代入する
	p = &orange
	fmt.Println(p)
	fmt.Println(*p / 30)
	*p = 81
	fmt.Println(*p / 3)
}
