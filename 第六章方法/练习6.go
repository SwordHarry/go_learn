// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := IntSet{words: []uint64{}}
	s.Add(1)
	s.Add(2)
	fmt.Println(s.Len())
	fmt.Println(s.String())
	s.Remove(2)
	fmt.Println(s.String())
	s2 := s.Copy()
	s2.Add(4)
	s2.Add(3)
	//s.Clear()
	fmt.Println(s2.String(), s.String())
	//s2.Intersection(&s)
	//fmt.Println(s2.String())
	s2.Different(&s)
	fmt.Println(s2.String())
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(val ...int) {
	for _, val := range val {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func (s *IntSet) Intersection(t *IntSet) {
	words := make([]uint64, min(len(s.words), len(t.words)))
	for i, tword := range t.words {
		if i < len(words) {
			words[i] = s.words[i] & tword
		} else {
			break
		}
	}
	s.words = words
}

// 差集
func (s *IntSet) Different(t *IntSet) {
	for i, word := range t.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			curWord := 64*i + j
			if word&(1<<uint(j)) != 0 && s.Has(curWord) {
				s.Remove(curWord)
			}
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) Len() int {
	result := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result++
			}
		}
	}
	return result
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	result := IntSet{words: []uint64{}}
	for _, word := range s.words {
		result.words = append(result.words, word)
	}
	return &result
}
