package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice % + v\n", mySlice)

}
