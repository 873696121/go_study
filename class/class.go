package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) Show() {
	fmt.Println(this.Name, this.Ad, this.Level)
}

func main() {
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("lisi")

	hero.Show()
}
