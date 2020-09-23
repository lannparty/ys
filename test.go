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
	case []interface{}:
		appendingMap := subset{}
		appendingMap[appendingKey] = appendingValue
		pointerType = append(pointerType, appendingMap)
		//fmt.Println(pointerType)
	case subset:
		fmt.Println("subset")
		pointerType[appendingKey] = appendingValue
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
	var emptyArray []interface{}
	fmt.Println("address", &emptyArray)
	//m3 := emptyArray
	//var emptyArray []interface{}
	m["test"] = m2
	m2["test2"] = &emptyArray
	//m3["test3"] = emptyArray
	fmt.Println(m)
	//appendWhole(m, "test3", "test4")
	emptyArray = append(emptyArray, "test")
	fmt.Println(*m2["test2"].([]interface{}))
}
