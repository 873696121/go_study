package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"yanyuan"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbaizhi"}}

	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Printf("%s\n", jsonStr)

	// 反序列化
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}

	fmt.Println(myMovie)
}
