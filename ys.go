package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[string]interface{}

func appendWholeItemToMap(targetMap subset, appendingItemKey string, appendingItemValue interface{}) {
        pointer := targetMap
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        pointer[appendingItemKey] = appendingItemValue
}

func appendNextItemToMap(targetMap subset, appendingItemKey string) {
        pointer := targetMap
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        pointer[appendingItemKey] = subset{}
}

func searchMap(child subset, cache subset, target string) {
        for key, _ := range child {
		nextCache := subset{}
		for cacheKey, cacheValue := range cache {
			nextCache[cacheKey] = cacheValue
		}
                appendNextItemToMap(nextCache, key)

                switch nextChild := child[key].(type) {
                case string:
			fmt.Println(nextChild)
                case subset:
			fmt.Println(nextChild)
                        searchMap(nextChild, nextCache, target)
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

	cache := subset{}
        appendWholeItemToMap(testMap3, "testMap123", testMap4)
        searchMap(m, cache, "us-west-2")
	//fmt.Println(m)

}
