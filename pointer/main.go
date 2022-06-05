package main

import (
	"fmt"
)

func main() {
	str := "a string value"
	pointer := &str        //& 是取地址符号
	anotherString := *&str //*是指针运算符  对于指针，必须使用 * 操作符来解引用
	fmt.Println(str)
	fmt.Println(pointer)       //取地址 0xc000010250
	fmt.Println(anotherString) //地址上面的值还是 a string value
	str = "changed string"
	fmt.Println(str)           //
	fmt.Println(pointer)       //地址还是 0xc000010250
	fmt.Println(anotherString) // 还是地址0xc000010250 上面的值
	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)
	changeParameter(&para, "bbb")
	fmt.Println(para)
	cannotChangeParameter(para, "ccc")
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}
