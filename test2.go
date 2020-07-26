package main

import "fmt"

type subset map[interface{}]interface{}

func main() {
	mList := []string{}
	m := subset{}
	m["test"] = 2
	m["test2"] = 3

	for key, _ := range m {
		mList = append(mList, key.(string))
	}
	fmt.Println(mList)
}
