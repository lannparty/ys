package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type subset map[interface{}]interface{}

// Append entire map to the end of target map.
func appendWhole(target subset, appendingKey interface{}, appendingValue interface{}) {
	var pointer interface{}
	pointer = target
	for len(pointer.(subset)) != 0 {
		for key, _ := range pointer.(subset) {
			switch pointerType := pointer.(subset)[key].(type) {
			case []interface{}:
				pointer = pointerType[0]
			case interface{}:
				pointer = pointer.(subset)[key].(subset)
			}
		}
	}
	//fmt.Println("pointer", pointer)
	//fmt.Println("Key", appendingKey, appendingValue)
	//pointer[appendingKey] = appendingValue
}

func main() {
	var content []byte
	reader := bufio.NewReader(os.Stdin)
	buffer := new(strings.Builder)
	_, _ = io.Copy(buffer, reader)

	content = []byte(buffer.String())

	unmarshalledContent := subset{}
	err := yaml.Unmarshal(content, unmarshalledContent)
	if err != nil {
		log.Fatal(err)
	}

	appendWhole(unmarshalledContent, "test", "value")

}
