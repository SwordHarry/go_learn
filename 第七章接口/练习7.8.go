// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type multiTrack struct {
	t      []*Track
	first  string
	second string
}

func (m *multiTrack) Len() int {
	return len(m.t)
}

func (m *multiTrack) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

func (m *multiTrack) Less(i, j int) bool {
	switch m.first {
	case "Title":
		if m.t[i].Title != m.t[j].Title {
			return m.t[i].Title < m.t[j].Title
		} else {
			return switchTrack(m, m.second, i, j)
		}
	case "Artist":
		if m.t[i].Artist != m.t[j].Artist {
			return m.t[i].Artist < m.t[j].Artist
		} else {
			return switchTrack(m, m.second, i, j)
		}
	case "Album":
		if m.t[i].Album != m.t[j].Album {
			return m.t[i].Album < m.t[j].Album
		} else {
			return switchTrack(m, m.second, i, j)
		}
	case "Year":
		if m.t[i].Year != m.t[j].Year {
			return m.t[i].Year < m.t[j].Year
		} else {
			return switchTrack(m, m.second, i, j)
		}
	case "Length":
		if m.t[i].Length != m.t[j].Length {
			return m.t[i].Length < m.t[j].Length
		} else {
			return switchTrack(m, m.second, i, j)
		}
	default:
		return m.t[i].Title < m.t[j].Title
	}
}

func switchTrack(m *multiTrack, caseStr string, i int, j int) bool {
	switch caseStr {
	case "Title":
		return m.t[i].Title < m.t[j].Title
	case "Artist":
		return m.t[i].Artist < m.t[j].Artist
	case "Album":
		return m.t[i].Album < m.t[j].Album
	case "Year":
		return m.t[i].Year < m.t[j].Year
	case "Length":
		return m.t[i].Length < m.t[j].Length
	default:
		return m.t[i].Title < m.t[j].Title
	}
}

func newMultiTrack(t []*Track, first string, second string) *multiTrack {
	return &multiTrack{
		t, first, second,
	}
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

func main() {
	m := newMultiTrack(tracks, "Title", "Length")
	sort.Sort(m)
	printTracks(m.t)
}
