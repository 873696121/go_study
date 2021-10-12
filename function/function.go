package main

import "fmt"

func foo1(a string, b string) int {
	fmt.Println(a)
	fmt.Println(b)
	return 100
}

// 返回多个返回值， 匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 66, 77
}

// 返回多个返回值， 有命名
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 100
	r2 = 200
	return
}

func main() {
	c := foo1("abc", "233")
	fmt.Println(c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1: ", ret1, "ret2: ", ret2)

	ret1, ret2 = foo2("haha", 999)
	fmt.Println("ret1: ", ret1, "ret2: ", ret2)
}
