//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("%v\tfor url:%v\n", err, url)
			os.Exit(1)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("%v\tfor url:%v\n", err, url)
			os.Exit(2)
		}
		fmt.Printf("%v\n", string(body))
	}
}
