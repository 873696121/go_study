package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	var slice2 []int

	slice2 = make([]int, 3)

	slice3 := make([]int, 4)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
