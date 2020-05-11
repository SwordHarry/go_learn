package main

import (
	"fmt"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

const (
	notVisit = iota
	visiting
	visited
)

//!+main
// func main() {
// 	for i, course := range topoSort(prereqs) {
// 		fmt.Printf("%d:\t%s\n", i+1, course)
// 	}
// }

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]int)
	var visitAll func(items []string)
	var hasRing bool
	visitAll = func(items []string) {
		for _, item := range items {
			if hasRing {
				return
			}
			if seen[item] == notVisit {
				seen[item] = visiting
				visitAll(m[item])
				seen[item] = visited
				order = append(order, item)
			} else if seen[item] == visiting {
				fmt.Println("存在环")
				hasRing = true
				return
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
