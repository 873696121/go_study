package main

import "fmt"

type Human struct {
	Name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("human eat...")
}

func (this *Human) Walk() {
	fmt.Println("human walk...")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("superman eat...")
}

func (this *SuperMan) Walk() {
	fmt.Println("superman walk...")
}

func main() {
	h := Human{Name: "zhangsan", sex: "male"}
	fmt.Println(h)
	h.Eat()
	h.Walk()

	s := SuperMan{Human{"li4", "femail"}, 88}
	fmt.Println(s)
	s.Eat()
	s.Walk()
}
