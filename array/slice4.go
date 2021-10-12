package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)
}
