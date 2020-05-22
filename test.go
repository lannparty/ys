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
	//testMap := subset{}
	//testMap["key"] = "value"
	//testMap["key2"] = "value2"
	//testMap["key3"] = "value3"

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
	//testMap := subset{}
	//testMap["key"] = "value"
	//testMapCopy := testMap
	//fmt.Println(testMap)
	//fmt.Println(testMapCopy)
	//testMapCopy["key2"] = "value2"
	//fmt.Println(testMap)
	//fmt.Println(testMapCopy)
	//var test interface{} = "test"
	//switch testSwitch := test.(type) {
	//case string:
	//	fmt.Println("string")
	//case int:
	//	fmt.Println("int")
	//}
	//testCopy := copyMap(testMap)
	//fmt.Println(testCopy)
	//fmt.Println(testMap)
	//testCopy["hello"] = "world"
	//fmt.Println(testCopy)
	//fmt.Println(testMap)
	//for key, value := range testMap {
	//	someCache := copyMap(testMap)
	//	fmt.Println(key)
	//	fmt.Println(value)
	//	someCache["key4"] = value
	//	fmt.Println(someCache)
	//}
        list := []int{1, 2, 3, 4}
        sum := 0
        for _ = range list {
                fmt.Println(add(1, sum))
        }

}
func add(i int, sum int) int {
        return i + sum
}

//func mod(t test) {
//	t["hello"] = "world"
//}

//func copyMap(target subset) subset {
//	targetCopy := subset{}
//	for key, value := range target {
//		targetCopy[key] = value
//	}
//	return targetCopy
//}

//func appendNextItemToMap(targetMap interface{}, appendingItemKey interface{}) {
//        pointer := targetMap
//        for len(pointer.(subset)) != 0 {
//                for _, value := range pointer.(subset) {
//                        //pointer = value.(subset)
//                        pointer = value
//                }
//        }
//        //pointer[appendingItemKey] = subset{}
//        //switch appendingType := appendingItemKey.(type) {
//        //        case string:
//        //                pointer[appendingItemKey] = appendingItemKey
//        //        case interface{}:
//        //                pointer[appendingItemKey] = subset{}
//        //}
//        switch appendingType := appendingItemKey.(type) {
//                case interface{}:
//                        pointer.(subset)[appendingItemKey] = subset{}
//                case string:
//                        //fmt.Println("string", appendingType)
//                        pointer.(subset)[appendingItemKey] = appendingType
//                case nil:
//                        pointer.(subset)[appendingItemKey] = nil
//        }
//}
