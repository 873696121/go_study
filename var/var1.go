package main

import "fmt"

var gA int = 100
var gB = 100

func main() {
	var a int
	fmt.Print(a)
	fmt.Printf("  %T\n", a)

	var b int = 100
	fmt.Print(b)
	fmt.Printf("  %T\n", b)

	var c = 100
	fmt.Print(c)
	fmt.Printf("  %T\n", c)

	d := 100
	fmt.Print(d)
	fmt.Printf("  %T\n", d)

	e := "abcd"
	fmt.Print(e)
	fmt.Printf("  %T\n", e)

	var x, y int = 100, 101
	fmt.Println("x == ", x, "y == ", y)
}
