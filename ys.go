package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var version string

var (
	read    string
	desired string

	rootCmd = &cobra.Command{
		Use:   "ys",
		Short: "Search yaml file.",
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			run(read, desired)
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ys",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ys-%s\n", version)
		},
	}
)

func Execute(v string) error {
	version = v
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&read, "read", "r", "", "YAML file to read")
	rootCmd.AddCommand(versionCmd)
}

func run(read string, desired string) {
	content, err := ioutil.ReadFile(read)
	if err != nil {
		log.Fatal(err)
	}
	cache := subset{}
	unmarshalledContent := subset{}
	err = yaml.Unmarshal(content, unmarshalledContent)
	fmt.Println(unmarshalledContent)
	printPathToDesired(unmarshalledContent, cache, "us-west-2")
}

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
			printPathToDesired(nextTarget, nextCache, desired)
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
			printPathToDesiredAndChildren(nextTarget, nextCache, desired)
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
			printDesiredAndChildren(nextTarget, desired)
		}
	}
}

func main() {
	Execute(version)
}
