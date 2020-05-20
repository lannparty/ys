package main

import (
        "fmt"
	//"io/ioutil"
	//"log"
	//"gopkg.in/yaml.v2"

)

type subset map[string]interface{}

func main() {
	//testMap := test{}
	//testMap2 := test{}
	//testMap3 := test{}
	//testMap4 := test{}
	//testMap5 := test{}
	//testMap4["key4"] = testMap5
	//testMap3["key3"] = testMap4
	//testMap2["key2"] = testMap3
	//testMap["key"] = testMap2

	//content, err := ioutil.ReadFile("test.yaml")
        //if err != nil {
        //        log.Fatal(err)
        //}

        //m := subset{}
        //err = yaml.Unmarshal(content, m)
        //if err != nil {
        //        log.Fatalf("cannot unmarshal data: %v", err)
        //}

	//fmt.Println(m)

	//fmt.Println(testMap)
	//fmt.Println("pointer", pointer)
	//mod(pointer)
	//fmt.Println(testMap)
	testMap := subset{}
	testMap["key"] = "value"
	testMapCopy := testMap
	fmt.Println(testMap)
	fmt.Println(testMapCopy)
	testMapCopy["key2"] = "value2"
	fmt.Println(testMap)
	fmt.Println(testMapCopy)
}

//func mod(t test) {
//	t["hello"] = "world"
//}
