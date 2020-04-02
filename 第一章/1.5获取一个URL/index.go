package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const HTTP string = "http://"
	for _, url := range os.Args[1:] {
		// 练习1.8 添加 http 协议头
		if !strings.HasPrefix(url, HTTP) {
			url = HTTP + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// var b io.Writer
		_, err = io.Copy(os.Stdout, resp.Body) // 练习1.7 赋值到标准输出
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// 练习 1.9 输出状态码
		fmt.Printf("%s", resp.Status)
	}
}
