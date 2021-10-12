package main

import "fmt"

func printArray(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

func main() {
	myArray := []int{1, 2, 3, 4}
	printArray(myArray)
}
