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
func appendWhole(target subset, appendingKey interface{}, appendingValue interface{}) {
        pointer := target
        for len(pointer) != 0 {
                for key, _ := range pointer {
                        pointer = pointer[key].(subset)
                }
        }
        pointer[appendingKey] = appendingValue
}

// Make a nil map in target with key appendingItemKey.
func appendNext(target interface{}, appendingItemKey interface{}) {
        pointer := target
        for len(pointer.(subset)) != 0 {
                for _, value := range pointer.(subset) {
                        pointer = value
                }
        }
        pointer.(subset)[appendingItemKey] = subset{}
}

// Marshal map and print.
func marshalledprint(target interface{}) {
        marshalledTarget, err := yaml.Marshal(target)
	if err != nil {
                log.Fatal(err)
	}
        fmt.Println(string(marshalledTarget))
}

// Print path to desired.
func printPathToDesired(target interface{}, cache subset, desired string) {
        for key, _ := range target.(subset) {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[key].(type) {
		case string:
			appendWhole(nextCache, key, nextTarget)
                        if nextTarget == desired {
                                marshalledprint(nextCache)
                        }
		case interface{}:
                        if key.(string) == desired {
                                printingCache := copyMap(nextCache)
		                appendNext(printingCache, key)
                                marshalledprint(printingCache)
                        }
		        appendNext(nextCache, key)
		        searchMap(nextTarget, nextCache, desired)
		case nil:
		        appendNext(nextCache, key)
                        if key.(string) == desired {
                                marshalledprint(nextCache)
                        }
		}
        }
}

// Print path to desired and children of desired.
func printPathToDesiredAndChildren(target interface{}, cache subset, desired string) {
        for key, _ := range target.(subset) {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[key].(type) {
		case string:
			appendWhole(nextCache, key, nextTarget)
                        if nextTarget == desired {
                                marshalledprint(nextCache)
                        }
		case interface{}:
                        if key.(string) == desired {
                                printingCache := copyMap(nextCache)
			        appendWhole(printingCache, key, nextTarget)
                                marshalledprint(printingCache)
                        }
		        appendNext(nextCache, key)
		        searchMap(nextTarget, nextCache, desired)
		case nil:
		        appendNext(nextCache, key)
                        if key.(string) == desired {
                                marshalledprint(nextCache)
                        }
		}
        }
}

// Print desired and children of desired.
func printDesiredAndChildren(target interface{}, desired string) {
        for key, _ := range target.(subset) {
                  switch nextTarget := target.(subset)[key].(type) {
                  case string:
                        if nextTarget == desired {
                                marshalledprint(nextTarget)
                        }
                  case interface{}:
                        if key.(string) == desired {
                                desiredMap := subset{key: nextTarget}
                                marshalledprint(desiredMap)
                        }
                        searchTarget(nextTarget, desired)
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

	//cache := subset{}
        //searchMap(m, cache, "account")
        searchTarget(m, "account")
}
