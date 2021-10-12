package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
