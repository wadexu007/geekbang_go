package main

import "fmt"

//接口（interface）是一种类型，一种抽象的类型。interface是一组方法的集合，接口做的事情就像是定义一个协议（规则），不关心对方是什么类型，只关心对方能做什么

type Stringer interface {
	Str() string
}

type Article struct {
	Title  string
	Author string
}

type Book struct {
	Title  string
	Author string
	Page   int
}

func (a Article) Str() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func (b Book) Str() string {
	return fmt.Sprintf("The %q article was written by %s.", b.Title, b.Author)
}

func main() {
	a := Article{
		Title:  "Golang",
		Author: "Sammy Shark",
	}
	Print(a)

	b := Book{
		Title:  "Python",
		Author: "Jenny James",
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.Str())
}

// func Print(a Article) {
// 	fmt.Println(a.Str())
// }

// func Print(b Book) {
// 	fmt.Println(b.Str())
// }
