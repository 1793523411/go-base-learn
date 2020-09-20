package main

import "fmt"

// Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”
// !调用有返回值的函数时，可以不接收其返回值
//!参数
// 函数的参数中如果相邻变量的类型相同，则可以省略类型,可变参数是指函数的参数数量不固定。
// o语言中的可变参数通过在参数名后加...来标识,可变参数通常要作为函数的最后一个参数,固定参数搭配可变参数使用时，可变参数要放在固定参数的后面,本质上，函数的可变参数是通过切片来实现的。
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//!返回值
// Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return nil
}

func main() {
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
}
