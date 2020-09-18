package main

import (
	"fmt"
)

type subset map[interface{}]interface{}

func endCheck(target interface{}) bool {
	end := false
	switch targetType := target.(type) {
	case []interface{}:
		if len(targetType) == 0 {
			end = true
		}
	case subset:
		if len(targetType) == 0 {
			end = true
		}
	}
	return end
}

func appendWhole(target subset, appendingKey interface{}, appendingValue interface{}) {
	var pointer interface{}
	pointer = target
	for endCheck(pointer) == false {
		for _, value := range pointer.(subset) {
			switch valueType := value.(type) {
			case []interface{}:
				pointer = valueType
			case interface{}:
				pointer = valueType
			}
		}
	}
	switch pointerType := pointer.(type) {
	case subset:

	case []interface{}:
	}
}

func main() {
	//read := "test2.yaml"
	//content, err := ioutil.ReadFile(read)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(content))
	//
	//unmarshalledContent := subset{}
	//err = yaml.Unmarshal(content, unmarshalledContent)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(unmarshalledContent)
	//
	//appendWhole(unmarshalledContent, "test", "test2")

	m := subset{}
	m2 := subset{}
	m3 := subset{}
	//var emptyArray []interface{}
	m["test"] = m2
	m2["test2"] = m3
	//m3["test3"] = emptyArray
	fmt.Println(m)
	appendWhole(m, "test3", "test4")
	fmt.Println(m)
}
