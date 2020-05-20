package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[interface{}]interface{}

func appendToMap(target subset, item subset) {
        pointer := target
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        fmt.Println(item)
}

//func searchMap(child subset, cache subset, target interface{}) {
//        for key, _ := range child {
//                nextCache := cache
//                appendToMap(nextCache, key.(subset))
//
//                switch nextChild := child[key].(type) {
//                case string:
//                case map[interface{}]interface{}:
//                        searchMap(nextChild, nextCache, target)
//                }
//        }
//}

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

        cache := make(map[interface{}]interface{})
        cache2 := make(map[interface{}]interface{})
        cache["key3"] = cache2
        searchMap(m, cache, "gid-prod")

}
