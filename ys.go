package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[interface{}]interface{}

func copyMap(target subset) subset {
        targetCopy := subset{}
        for key, value := range target {
                targetCopy[key] = value
        }
        return targetCopy
}

func appendWholeItemToMap(targetMap subset, appendingItemKey interface{}, appendingItemValue interface{}) {
        pointer := targetMap
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        pointer[appendingItemKey] = appendingItemValue
}

func appendNextItemToMap(targetMap interface{}, appendingItemKey interface{}) {
        pointer := targetMap
        for len(pointer.(subset)) != 0 {
                for _, value := range pointer.(subset) {
                        //pointer = value.(subset)
                        pointer = value
                }
        }
        //pointer[appendingItemKey] = subset{}
        //switch appendingType := appendingItemKey.(type) {
        //        case string:
        //                pointer[appendingItemKey] = appendingItemKey
        //        case interface{}:
        //                pointer[appendingItemKey] = subset{}
        //}
        switch appendingType := appendingItemKey.(type) {
                case interface{}:
                        pointer.(subset)[appendingItemKey] = subset{}
                case string:
                        //fmt.Println("string", appendingType)
                        pointer.(subset)[appendingItemKey] = appendingType
                case nil:
                        pointer.(subset)[appendingItemKey] = nil
        }
}

//func searchMap(child interface{}, cache subset, target string) {
//        for key, _ := range child.(subset) {
//		nextCache := subset{}
//		for cacheKey, cacheValue := range cache {
//			nextCache[cacheKey] = cacheValue
//		}
//                appendNextItemToMap(nextCache, key)
//                fmt.Println(nextCache)
//
//                switch nextChild := child.(subset)[key].(type) {
//                case string:
//			//fmt.Println("string", cache)
//                case interface{}:
//			//fmt.Println("interface", cache)
//                        searchMap(nextChild, nextCache, target)
//                case nil:
//			//fmt.Println("nil", cache)
//                }
//        }
//}

func searchMap(child interface{}, cache subset, target string) {
        for key, _ := range child.(subset) {
		nextCache := copyMap(cache)
		appendNextItemToMap(nextCache, key)
                nextChild := child.(subset)[key].(subset)
                searchMap2(nextChild, nextCache, target)
                //fmt.Println("key", key)
		//fmt.Println(child.(subset)[key])
                //fmt.Println("nextCache", nextCache)
		//switch nextChild := child.(subset)[key].(type) {
		//case string:
                //        fmt.Println("string")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//	appendNextItemToMap(nextCache, nextChild)
		//case interface{}:
                //        fmt.Println("interface")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//        searchMap(nextChild, nextCache, target)
		//case nil:
                //        fmt.Println("nil")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//	appendNextItemToMap(nextCache, nextChild)
		//}
        }
}

func searchMap2(child interface{}, cache subset, target string) {
        for key, _ := range child.(subset) {
                //fmt.Println(key)
                //fmt.Println(cache)
		nextCache := copyMap(cache)
		appendNextItemToMap(nextCache, key)
                //fmt.Println(nextCache)
                //fmt.Println("nextCache2", nextCache)
                //fmt.Println("key", key)
		//fmt.Println(child.(subset)[key])
                //fmt.Println("nextCache", nextCache)
		//switch nextChild := child.(subset)[key].(type) {
		//case string:
                //        fmt.Println("string")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//	appendNextItemToMap(nextCache, nextChild)
		//case interface{}:
                //        fmt.Println("interface")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//        searchMap(nextChild, nextCache, target)
		//case nil:
                //        fmt.Println("nil")
                //        fmt.Println("nextCache", nextCache)
                //        fmt.Println("nextChild", nextChild)
		//	appendNextItemToMap(nextCache, nextChild)
		//}
        }

}

func main() {

	content, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	m := subset{}
	err = yaml.Unmarshal(content, m)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	testMap := subset{}
        testMap2 := subset{}
        testMap3 := subset{}
        testMap4 := subset{}
        testMap5 := subset{}
        testMap6 := subset{}
        testMap5["key5"] = testMap6
        testMap4["key4"] = testMap5

        testMap2["key2"] = testMap3
        testMap["key"] = testMap2
        testMap2["keycopy"] = "testMapCopy"
        //fmt.Println(testMap)

	cache := subset{}
        searchMap(m, cache, "us-west-2")
        copyCacheTest(testMap, cache)
	//fmt.Println(m)

}

func copyCacheTest(someMap interface{}, cache subset) {
        for key, _ := range someMap.(subset) {
                nextCache := copyMap(cache)
                appendNextItemToMap(nextCache, key)
                nextChild := someMap.(subset)[key].(subset)
                copyCacheTest2(nextChild, cache)
        }
}
func copyCacheTest2(someMap interface{}, cache subset) {
        for key, _ := range someMap.(subset) {
                nextCache := copyMap(cache)
                appendNextItemToMap(nextCache, key)
                fmt.Println(nextCache)
        }
}
