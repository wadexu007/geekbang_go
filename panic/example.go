package main

import (
	"fmt"
	//"os"
)

var user = ""

func inita() {
	defer func() {
		fmt.Print("defer##\n")
	}()
	if user == "" {
		fmt.Print("@@@before panic\n")
		panic("no value for user\n")
		// fmt.Print("!!after panic\n") //panic在user=""时，打断了函数的执行，fmt.Print("!!after panic\n")没有执行
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
			b = true
		}
	}()
	//直到执行完所有函数的defer，退出程序, 下面不会执行
	f()
	fmt.Print("after the func run")
	return
}
func main() {
	throwsPanic(inita)
}
