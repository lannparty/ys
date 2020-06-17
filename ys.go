package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[interface{}]interface{}

// Create a copy of a map and all its nested maps.
func copyMap(target subset) subset {
	targetCopy := subset{}
	for key, value := range target {
                switch value := value.(type) {
                        case string:
		                targetCopy[key] = value
                        case subset:
                                targetCopy[key] = copyMap(value)
                }
	}
	return targetCopy
}

// Append entire map to the end of target map.
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
                        pointer = value
                }
        }
        pointer[appendingItemKey] = subset{}
        switch appendingType := appendingItemKey.(type) {
                case string:
                        pointer[appendingItemKey] = appendingItemKey
                case interface{}:
                        pointer[appendingItemKey] = subset{}
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

// Create a cache for each path and return path to target when target is encountered.
func searchMap(child interface{}, cache subset, target string) {
        for key, _ := range child.(subset) {
		nextCache := copyMap(cache)
		appendNextItemToMap(nextCache, key)
		switch nextChild := child.(subset)[key].(type) {
		case string:
			appendNextItemToMap(nextCache, nextChild)
		case interface{}:
		        searchMap(nextChild, nextCache, target)
		case nil:
			appendNextItemToMap(nextCache, nextChild)
		}
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

	cache := subset{}
        searchMap(m, cache, "us-west-2")
	//fmt.Println(m)

}
