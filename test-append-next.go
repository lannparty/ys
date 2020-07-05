package main
import (
        "fmt"
)

type subset map[interface{}]interface{}

//func copyMap(target subset) subset {
//	targetCopy := subset{}
//	for key, value := range target {
//                switch value := value.(type) {
//                        case string:
//		                targetCopy[key] = value
//                        case subset:
//                                targetCopy[key] = copyMap(value)
//                }
//	}
//	return targetCopy
//}


// Append next item to map small test.
func appendNextItemToMap(targetMap interface{}) {
        pointer := targetMap
        for len(pointer.(subset)) != 0 {
                for _, value := range pointer.(subset) {
                        pointer = value
                }
        }
        fmt.Println(pointer)
}

func main() {
	testMap := subset{}
        //testMap3 := subset{}
        testMap2 := subset{}
	testMap["key2"] = testMap2
        testMap2["key3"] = nil
        fmt.Println(testMap)
        appendNextItemToMap(testMap)
}
