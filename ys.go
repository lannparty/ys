package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
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
	multi   bool

	rootCmd = &cobra.Command{
		Use:   "ys",
		Short: "Search yaml file.",
		Args:  cobra.MaximumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			run(read, desired, mode, filter, multi)
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
	rootCmd.Flags().StringVarP(&filter, "filter", "f", "", "Comma delimited list of strings to filter path to target. (Does not work in childonly mode because there's no path printed.)")
	rootCmd.Flags().StringVarP(&mode, "mode", "m", "", "Alternate ways of returning path. (pathonly, childonly)")
	rootCmd.Flags().BoolVarP(&multi, "multi", "i", false, "Operate on multiple yaml blocks independently, new line delimited.")
	rootCmd.AddCommand(versionCmd)
}

func run(read string, desired string, mode string, filter string, multi bool) {
	var content []byte
	var contentList [][]byte
	var splitYaml []string
	var err error
	if read == "" {
		reader := bufio.NewReader(os.Stdin)
		buffer := new(strings.Builder)
		_, err = io.Copy(buffer, reader)
		if err != nil {
			log.Fatal(err)
		}
		if multi == true {
			splitYaml = strings.Split(buffer.String(), "\n\n")
			for _, value := range splitYaml {
				contentList = append(contentList, []byte(value))
			}
			for _, value := range contentList {
				search(value)
			}
		} else {
			content = []byte(buffer.String())
			if err != nil {
				log.Fatal(err)
			}
			search(content)
		}
	} else {
		content, err = ioutil.ReadFile(read)
		if err != nil {
			log.Fatal(err)
		}
		search(content)
	}
}

type subset map[interface{}]interface{}

func search(content []byte) {
	cache := subset{}
	unmarshalledContent := subset{}
	err := yaml.Unmarshal(content, unmarshalledContent)
	if err != nil {
		log.Fatal(err)
	}
	switch mode {
	case "":
		printPathToDesiredAndChildren(unmarshalledContent, cache, desired, filter)
	case "pathonly":
		printPathToDesired(unmarshalledContent, cache, desired, filter)
	case "childonly":
		printDesiredAndChildren(unmarshalledContent, desired)
	}
}

// Loop through path and check if it contains filter.
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

// Split filter by comma delimited string and validateFilter for every filter.
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

// Combine two maps and all its nested maps.
func mergeMap(mainTarget subset, subTarget subset) subset {
	for key, value := range subTarget {
		switch value := value.(type) {
		case string:
			mainTarget[key] = value
		case subset:
			mainTarget[key] = copyMap(value)
		}
	}
	return mainTarget
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
func marshalledPrint(target interface{}) {
	marshalledTarget, err := yaml.Marshal(target)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshalledTarget))
}

// Print path to desired.
func printPathToDesired(target interface{}, cache subset, desired string, filter string) {
	var sortCache []string
	for key, _ := range target.(subset) {
		sortCache = append(sortCache, key.(string))
	}
	sort.Strings(sortCache)
	for _, value := range sortCache {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[value].(type) {
		case string:
			nextCacheCopy := copyMap(nextCache)
			appendNext(nextCacheCopy, value)
			appendWhole(nextCache, value, nextTarget)
			if nextTarget == desired && validateAllFilters(nextCacheCopy, filter) == true {
				marshalledPrint(nextCache)
			}
		case []interface{}:
			var emptyArray []interface{}
			appendWhole(nextCache, value, emptyArray)
			for _, value2 := range nextTarget {
				nextNextCache := copyMap(nextCache)
				switch valueType := value2.(type) {
				case string:
					nextCacheCopy := copyMap(nextCache)
					if valueType == desired && validateAllFilters(nextCacheCopy, filter) {
						appendWhole(nextCacheCopy, value, valueType)
						marshalledPrint(nextCacheCopy)
					}
				case subset:
					fmt.Println(value2)
					fmt.Println(reflect.TypeOf(value2))
				}
			}
		case interface{}:
			if value == desired && validateAllFilters(nextCache, filter) == true {
				printingCache := copyMap(nextCache)
				appendNext(printingCache, value)
				marshalledPrint(printingCache)
			}
			appendNext(nextCache, value)
			printPathToDesired(nextTarget, nextCache, desired, filter)
		case nil:
			appendNext(nextCache, value)
			if value == desired && validateAllFilters(nextCache, filter) == true {
				marshalledPrint(nextCache)
			}
		}
	}
}

// Print path to desired and children of desired.
func printPathToDesiredAndChildren(target interface{}, cache subset, desired string, filter string) {
	var sortCache []string
	for key, _ := range target.(subset) {
		sortCache = append(sortCache, key.(string))
	}
	sort.Strings(sortCache)
	for _, value := range sortCache {
		nextCache := copyMap(cache)
		switch nextTarget := target.(subset)[value].(type) {
		case string:
			if value == desired && validateAllFilters(nextCache, filter) == true {
				printingCache := copyMap(nextCache)
				appendWhole(printingCache, value, nextTarget)
				marshalledPrint(printingCache)
			}
			nextCacheCopy := copyMap(nextCache)
			appendNext(nextCacheCopy, value)
			appendWhole(nextCache, value, nextTarget)
			if nextTarget == desired && validateAllFilters(nextCacheCopy, filter) == true {
				marshalledPrint(nextCache)
			}
		case interface{}:
			if value == desired && validateAllFilters(nextCache, filter) == true {
				printingCache := copyMap(nextCache)
				appendWhole(printingCache, value, nextTarget)
				marshalledPrint(printingCache)
			}
			appendNext(nextCache, value)
			printPathToDesiredAndChildren(nextTarget, nextCache, desired, filter)
		case nil:
			appendNext(nextCache, value)
			if value == desired && validateAllFilters(nextCache, filter) == true {
				marshalledPrint(nextCache)
			}
		}
	}
}

// Print desired and children of desired.
func printDesiredAndChildren(target interface{}, desired string) {
	var sortCache []string
	for key, _ := range target.(subset) {
		sortCache = append(sortCache, key.(string))
	}
	sort.Strings(sortCache)
	for _, value := range sortCache {
		switch nextTarget := target.(subset)[value].(type) {
		case string:
			if nextTarget == desired {
				marshalledPrint(nextTarget)
			}
		case interface{}:
			if value == desired {
				desiredMap := subset{value: nextTarget}
				marshalledPrint(desiredMap)
			}
			printDesiredAndChildren(nextTarget, desired)
		}
	}
}

func main() {
	Execute(version)
}
