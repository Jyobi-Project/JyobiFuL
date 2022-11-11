package main

import "fmt"

// structの定義 -> structはフィールドの集まりらしいです。しらんけど
type fruit struct {
	x string
	y int
}

func main() {
	// fruitの型に合わせて値を代入して表示
	fmt.Println(fruit{"apple", 3})
}
