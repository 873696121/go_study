package main

import "fmt"

func main() {
	defer fmt.Println("main1 end")
	defer fmt.Println("main2 end")
	fmt.Println("111")
	fmt.Println("222")
}
