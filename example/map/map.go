package main

import (
	"fmt"
)

func main() {
	myMap := map[string]string{}
	myMap["a"] = "123"
	myMap["b"] = "345"
	fmt.Println(myMap["a"] + myMap["b"])

	myMap2 := map[string]string{
		"alicloudprovider": "alibaba",
		"jdcloudprovider":  "jd",
		"gcpcloudprovider": "google",
	}
	fmt.Println("cloudmap:", myMap2)

	for k, v := range myMap2 {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}
	myFunctionMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFunctionMap)
	f := myFunctionMap["funcA"]
	fmt.Println(f())

	for k, v := range myFunctionMap {
		fmt.Println("key:", k)
		fmt.Println("value:", v())
	}

	myFunctionMap1 := map[string]int{"sum": sum(2, 2)}
	fmt.Println(myFunctionMap1["sum"])
}

func sum(x, y int) int {
	println(x + y)
	return x + y
}
