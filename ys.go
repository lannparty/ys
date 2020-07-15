package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var version string

var (
	read    string
	filter  string
	desired string
	mode    string

	rootCmd = &cobra.Command{
		Use:   "ys",
		Short: "Search yaml file.",
		Args:  cobra.MaximumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			run(read, desired, mode, filter)
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
	rootCmd.Flags().StringVarP(&read, "read", "r", "", "YAML file to read.")
	rootCmd.Flags().StringVarP(&desired, "desired", "d", "", "Target subset to search for.")
	rootCmd.Flags().StringVarP(&filter, "filter", "f", "", "Comma delimited list of strings to filter output.")
	rootCmd.Flags().StringVarP(&mode, "mode", "m", "", "Alternate ways of returning path.")
	rootCmd.AddCommand(versionCmd)
}

func run(read string, desired string, mode string, filter string) {
	content, err := ioutil.ReadFile(read)
	if err != nil {
		log.Fatal(err)
	}
	cache := subset{}
	unmarshalledContent := subset{}
	err = yaml.Unmarshal(content, unmarshalledContent)
	switch mode {
	case "":
		printPathToDesiredAndChildren(unmarshalledContent, cache, desired, filter)
	case "pathonly":
		printPathToDesired(unmarshalledContent, cache, desired, filter)
	case "childonly":
		printDesiredAndChildren(unmarshalledContent, desired)
	}
}

type subset map[interface{}]interface{}

func validateFilter(target subset, filter string) bool {
	pointer := target
	var contains bool = false
	for len(pointer) != 0 {
		for key2, _ := range pointer {
			pointer = pointer[key2].(subset)
			if key2 == filter {
				contains = true
			}
		}
	}
	return contains
}

func validateAllFilters(target subset, filter string) bool {
	if filter == "" {
		return true
	}
	contains := []bool{}
	containsAll := true
	for _, value := range strings.Split(filter, ",") {
		contains = append(contains, validateFilter(target, value))
	}
	for _, value := range contains {
		if value == false {
			containsAll = false
		}
	}
	return containsAll
}

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
func printPathToDesired(target interface{}, cache subset, desired string, filter string) {
	for key, _ := range target.(subset) {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[key].(type) {
		case string:
			nextCacheCopy := copyMap(nextCache)
			appendNext(nextCacheCopy, key)
			appendWhole(nextCache, key, nextTarget)
			if nextTarget == desired && validateAllFilters(nextCacheCopy, filter) == true {
				marshalledprint(nextCache)
			}
		case interface{}:
			if key.(string) == desired && validateAllFilters(nextCache, filter) == true {
				printingCache := copyMap(nextCache)
				appendNext(printingCache, key)
				marshalledprint(printingCache)
			}
			appendNext(nextCache, key)
			printPathToDesired(nextTarget, nextCache, desired, filter)
		case nil:
			appendNext(nextCache, key)
			if key.(string) == desired && validateAllFilters(nextCache, filter) == true {
				marshalledprint(nextCache)
			}
		}
	}
}

// Print path to desired and children of desired.
func printPathToDesiredAndChildren(target interface{}, cache subset, desired string, filter string) {
	for key, _ := range target.(subset) {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[key].(type) {
		case string:
			nextCacheCopy := copyMap(nextCache)
			appendNext(nextCacheCopy, key)
			appendWhole(nextCache, key, nextTarget)
			if nextTarget == desired && validateAllFilters(nextCacheCopy, filter) == true {
				marshalledprint(nextCache)
			}
		case interface{}:
			if key.(string) == desired && validateAllFilters(nextCache, filter) == true {
				printingCache := copyMap(nextCache)
				appendWhole(printingCache, key, nextTarget)
				marshalledprint(printingCache)
			}
			appendNext(nextCache, key)
			printPathToDesiredAndChildren(nextTarget, nextCache, desired, filter)
		case nil:
			appendNext(nextCache, key)
			if key.(string) == desired && validateAllFilters(nextCache, filter) == true {
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
