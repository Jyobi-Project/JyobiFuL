package main

import "fmt"

// 戻り値に変数名を付けることが出来る -> 戻り値の意味を示す名前付けを推奨
func split(sum int) (apple, orange int) {
	apple = sum * 4 / 9
	orange = sum - apple
	return
}

func main() {
	fmt.Println(split(17))
}
