package main

import (
	"golang.org/x/net/html"
)

const url = "http://www.baidu.com"

// func main() {

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 		os.Exit(1)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		resp.Body.Close()
// 		os.Exit(1)
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	resp.Body.Close()
// 	m := map[string]int{}
// 	visit(&m, doc)
// 	fmt.Println(m)
// }

func visit(m *map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		(*m)[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(m, c)
	}
}
