package main

import "fmt"

func main() {
	//回调函数没有函数名，所以回调函数需要赋值到某个变量，或者作为立即执行的函数
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(1, 2)

	//立即执行的函数
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)
}
