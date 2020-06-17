package main
import (
        "fmt"
)

type subset map[string]interface{}

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

func main() {
        nestedMap := subset{}
        nestedMap["nestedKey"] = "nestedValue"
	testMap := subset{}
	testMap["key"] = "value"
	testMap["key2"] = "value2"
	testMap["key3"] = "value3"
        testMap["someNest"] = nestedMap
        testMapCopy := copyMap(testMap)
        fmt.Println(testMap)
        fmt.Println(testMapCopy)
	testMap["key4"] = "value4"
        nestedMap["nestedKey2"] = "nestedValue2"
        fmt.Println(testMap)
        fmt.Println(testMapCopy)
}
