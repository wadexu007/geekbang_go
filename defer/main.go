package main

import "fmt"

//规则一： 延迟函数的参数在defer语句出现时就已经确定
//规则二： 延迟函数执行按后进先出顺序执行， 即先出现的defer最后执行
//规则三： 延迟函数可能操作主函数的具名返回值

func d() int {
	var i int
	for i = 0; i < 4; i++ {
		defer fmt.Print(i, "\n")
	}
	return i
}

func main() {
	fmt.Println("result: ", d())
}

//结果输出的是0
//defer函数会在return之后被调用, 但这个变量(i)在defer被声明的时候，就已经确定其确定的值了
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//先输出1，再输出0
//defer执行顺序为先进后出
func b() {
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}

//defer可以读取有名返回值
//当执行return 1 之后，i的值就是1. 此时此刻，defer代码块开始执行，对i进行自增操作。 因此输出2.
func c() (i int) {
	defer func() { i++ }()
	return 1
}
