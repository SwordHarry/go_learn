package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const URL = "https://xkcd.com/571/info.0.json"

type Result struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Alt        string
	Img        string
	Title      string
	Day        string
	Transcript string
}

//!+template
const templ = `Month: {{.Month}}
Number: {{.Num}}
Link:   {{.Link}}
SafeTitle:  {{.SafeTitle}}
Day:    {{.Day }} days
`

//!-template

func SearchIssues() (*Result, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	// data, err := json.MarshalIndent(resp.Body, "", "   ")
	// if err != nil {
	// 	resp.Body.Close()
	// 	return nil, err
	// }
	resp.Body.Close()
	return &result, nil
}

var report = template.Must(template.New("issuelist").
	Parse(templ))

func main() {
	result, err := SearchIssues()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
