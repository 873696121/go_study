package main

import "fmt"

// 无缓冲
func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine ...")

		c <- 666
	}()

	num := <-c
	fmt.Println(num)
}
