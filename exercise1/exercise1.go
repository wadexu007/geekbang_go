package main

import (
	"fmt"
)

func main() {

	myArray := [5]string{"I", "am", "stupid", "and", "weak"}

	for index := range myArray {
		if myArray[index] == "stupid" {
			myArray[index] = "smart"
		}
		if myArray[index] == "weak" {
			myArray[index] = "strong"
		}
	}
	fmt.Println(myArray)
}
