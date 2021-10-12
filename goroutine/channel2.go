package main

import (
	"fmt"
	"time"
)

// 有缓冲
func main() {
	c := make(chan int, 3)

	fmt.Println("len, ", len(c), " cap,", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行， 发送的是 ", i, " len:", len(c), " cap:", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num, ok := <-c
		fmt.Println(num, " ", ok)
	}
	defer println("main结束")
}
