package main

import (
	"log"
	"os"
)

var tokens []string

func main() {
	if len(os.Args) > 1 {
		load_json(os.Args[1])
	}
}

func load_json(file string) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	json := string(bytes[:])
	Lex(json)
	Parse(tokens)
}
