package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type subset map[interface{}]interface{}

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

	test := m["provider"].(subset)["Azure"].(subset)["account"]
        fmt.Println(test)
        test = "bob"
        fmt.Println(m)

}
