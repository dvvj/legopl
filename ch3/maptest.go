package main

import (
	"fmt"
	"sort"
)

func sortedPrint(m map[string]int) {
	//var names []string
	names := make([]string, 0, len(m))

	for name := range m {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s: %d\n", name, m[name])
	}
}

func main() {
	ages := map[string]int{
		"bob":   12,
		"alice": 14,
		"cindy": 11,
	}

	age, ok := ages["diana"]
	if !ok {
		fmt.Printf("diana not found, result: %d\n", age)
	} else {
		fmt.Printf("diana: %d\n", age)
	}

	sortedPrint(ages)
}
