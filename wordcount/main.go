//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processFile(filepath string, hashMap map[string]int) (err error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("supposingly, wrong filePath: \"%v\":%w", filepath, err)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		hashMap[line]++
	}
	return nil
}

func main() {
	filePaths := os.Args[1:]

	lineCount := map[string]int{}

	for _, file := range filePaths {
		err := processFile(file, lineCount)
		check(err)
	}
	for k, v := range lineCount {
		if v > 1 {
			fmt.Printf("%v\t%v\n", v, k)
		}
	}
}
