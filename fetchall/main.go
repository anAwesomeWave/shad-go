//go:build !solution

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(url string, c chan struct{}) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		log.Printf("%v err: %v", url, err)
	}
	log.Printf("%v\t%v\n", time.Since(start), url)
	c <- struct{}{}
}

func main() {
	urls := os.Args[1:]
	start := time.Now()
	c := make(chan struct{})

	for _, url := range urls {
		go fetch(url, c)
	}
	for range urls {
		<-c
	}
	fmt.Printf("Program ended with total time: %v", time.Since(start))
}
