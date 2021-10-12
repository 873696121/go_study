package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called...")
	fmt.Println(this)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println(inputType)

	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Type, " ", m.Name)
	}
}

func main() {
	num := 1.2345
	reflectNum(num)

	user := User{1, "Huhong", 16}
	DoFiledAndMethod(user)
}
