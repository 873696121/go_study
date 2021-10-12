package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "beijing"
	cityMap["Japan"] = "tokyo"
	cityMap["USA"] = "NewYork"

	for key, value := range cityMap {
		fmt.Println("key : ", key, " value : ", value)
	}

	delete(cityMap, "China")
	fmt.Println(cityMap)

	cityMap["USA"] = "DC"
	fmt.Println(cityMap)

}
