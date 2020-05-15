package main

import (
	"fmt"
	"sort"
	"strings"
)

func sqrt(n int) int { return n * n }
func neg(n int) int  { return -n }

func test1() {
	var f func(int) int

	//fmt.Println(f(20)) // panic

	f = sqrt
	fmt.Println(f(20))
	f = neg
	fmt.Println(f(20))
}

func add1(r rune) rune { return r + 1 }

func testStringsMap() {
	fmt.Println(strings.Map(add1, "abc"))
	fmt.Println(strings.Map(add1, "123"))
	fmt.Println(strings.Map(add1, "你好"))
}

func main() {
	// test1()
	// testStringsMap()

	testTopoSort()
}

const Intro = "intro"
const Algorithm = "algo"
const DataStructure = "data structure"

func testTopoSort() {
	var prereqs = map[string][]string{
		DataStructure: {Intro},
		Algorithm:     {DataStructure},
	}

	orders := topoSort(prereqs)
	for i, item := range orders {
		fmt.Printf("%d:\t%s\n", i+1, item)
	}
}

func topoSort(prereq map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(prereq[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range prereq {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
