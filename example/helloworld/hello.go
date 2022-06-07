package main

import (
	"fmt"
)

func main() {
	name := "test"
	fmt.Printf("%s\n", name)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fullString := "hello world"
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}
}
