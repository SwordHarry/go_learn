package main

import "net/http"

import "log"

import "fmt"

func main() {
	http.HandleFunc("/", handler) // 回调处理程序
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
