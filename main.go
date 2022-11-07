package main

import (
	"fmt"
	"os"
	"runtime"
)

var EOL = "\n"
var content string

func main() {
	if runtime.GOOS == "windows" {
		EOL = "\r\n"
	}
	if len(os.Args) > 1 {
		bytes, err := os.ReadFile(os.Args[1])
		handle(err)
		content = string(bytes)
		parse()
		// fmt.Println(tokens)
		fmt.Printf("\n")
		for _, tok := range tokens {
			fmt.Println(tok)
		}
	} else {
		fmt.Println("USAGE: jsonp [FILE]")
	}
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
