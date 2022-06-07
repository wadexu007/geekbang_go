package main

import "fmt"

type animal struct {
	Name  string
	Color string
	Age   int
}

//方法是作用，用在指定接受者，身上的函数
func (a *animal) Run() {
	fmt.Println(a.Name, "在奔跑")
}

func (a *animal) jump() {
	fmt.Println(a.Name, "跳")
}

type Cat struct {
	Cat animal
}

func main() {
	cat1 := Cat{
		Cat: animal{
			Name:  "Tom",
			Color: "Black",
			Age:   10,
		},
	}
	fmt.Println(cat1.Cat.Name)
	cat1.Cat.Run()
}
