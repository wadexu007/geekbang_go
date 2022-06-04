package main

import "fmt"

func b() int {
	var i int
	for i = 0; i < 4; i++ {
		defer fmt.Print(i, "\n")
	}
	return i
}

func main() {
	fmt.Println("result: ", b())
}
