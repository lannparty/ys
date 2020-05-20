package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[string]interface{}

func appendToMap(targetMap subset, appendingItemKey string, appendingItemValue interface{}) {
        pointer := targetMap
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        pointer[appendingItemKey] = appendingItemValue
}

func searchMap(child subset, cache subset, target interface{}) {
        for key, _ := range child {
                nextCache := cache
                appendToMap(nextCache, key.(subset))

                switch nextChild := child[key].(type) {
                case string:
                case map[interface{}]interface{}:
                        searchMap(nextChild, nextCache, target)
                }
        }
}

func main() {

	content, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[interface{}]interface{})
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

        //searchMap(m, cache, "gid-prod")
        appendToMap(testMap3, "testMap123", testMap4)
	fmt.Println(testMap)

}
