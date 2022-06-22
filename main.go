package main

import (
  "os"
  "log"
  lexer "json_parser/lexer"
  parser "json_parser/parser"
)

func main()  {
  if len(os.Args) > 1 {
    load_json(os.Args[1])
  }
}

func load_json(file string) {
  bytes, error := os.ReadFile(file)
  if error != nil {
    log.Fatal(error)
  }
  json := string(bytes[:])
  lexer.Init(json)
  parser.Init(lexer.Tokens)
}
