package main

//反射是指在程序运行期对程序本身进行访问和修改的能力，Go程序在运行期使用reflect包访问程序的反射信息。

import (
	"fmt"
	"reflect"
)

type Member struct {
	id   int
	name string
	age  int
}

func main() {

	member := Member{1, "Adam", 100}

	t := reflect.TypeOf(member)  //取得所有元素
	v := reflect.ValueOf(member) //取得值

	fmt.Println(t) //output main.Member
	fmt.Println(v) //output {1 Adam 100}

}
