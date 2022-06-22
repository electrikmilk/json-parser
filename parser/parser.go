package parser

import (
  "fmt"
  "strings"
  "log"
)

var parse_tokens []string

func Init(tokens []string) {
  parse_tokens = tokens
  var last_token string
  for i, t := range tokens {
    fmt.Println(i,t)

    /* GRAMMAR CHECKS */

    // strings
    if strings.Contains(t,"STRING") && last_token != "LQUOTE" {
      log.Fatal("Invalid string")
    }
    if strings.Contains(t,"RQUOTE") && last_token != "STRING" {
      log.Fatal("Invalid syntax, ending quote with no string")
    }
    if strings.Contains(t,"LQUOTE") && (last_token != "COLON" || last_token != "EOL") {
      log.Fatal("Invalid syntax, starting quote with no colon or after new line")
    }
    if t == "RQUOTE" && next_token(i) != "COMMA" && next_token(i) != "COLON" {
      log.Fatal("Failed to seperate with comma")
    }

    last_token = strings.Split(t,":")[0]
  }
}

func next_token(index int) string {
  index += 1
  return strings.Split(parse_tokens[index],":")[0]
}
