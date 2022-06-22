package lexer

import (
  "fmt"
  "strings"
)

var c string
var pos int
var line_chars []string
var Tokens []string

const digits string = "0123456789"

func Init(json string) {
  line_chars = strings.Split(json,"")
  c = line_chars[pos]
  for len(c) > 0 {
    if c != "\t" && c != " " {
      if c == "\n" {
        if(len(line_chars) == (pos+1)) {
          Tokens = append(Tokens, "EOF")
        } else {
          Tokens = append(Tokens, "EOL")
        }
      } else if c == "{" {
        Tokens = append(Tokens, "LCURLY")
      } else if c == "}" {
        Tokens = append(Tokens, "RCURLY")
      } else if c == ":" {
        Tokens = append(Tokens, "COLON")
      } else if c == "\"" {
        Tokens = append(Tokens, "LQUOTE")
        advance()
        token_string()
      } else if c == "[" {
        Tokens = append(Tokens, "LBRACKET")
      } else if c == "]" {
        Tokens = append(Tokens, "RBRACKET")
      } else if c == "," {
        Tokens = append(Tokens, "COMMA")
      } else if c == "t" || c == "f" {
        token_bool()
      } else if strings.Contains(digits,c) {
        token_number()
      }
    }
    advance()
  }
}

func token_number() {
  var number string
  for strings.Contains(digits,c) {
    number = number + c
    advance()
  }
  number = fmt.Sprintf("INT:%s",number)
  Tokens = append(Tokens,number)
}

func token_string() {
  var str string
  for c != "\"" {
      str = str + c
      advance()
  }
  str = fmt.Sprintf("STRING:%s",str)
  Tokens = append(Tokens,str,"RQUOTE")
}

func token_bool() {
  var peek string = peek(4)
  if peek == "true" {
    Tokens = append(Tokens,"BOOL:true")
  } else if peek == "false" {
    Tokens = append(Tokens,"BOOL:false")
  }
}

func advance() {
  pos += 1;
  if pos < len(line_chars) {
    c = line_chars[pos]
  } else {
    c = ""
  }
}

func peek(ahead int) string {
  ahead = pos + ahead
  var future string
  for i := pos; i < ahead; i++ {
    future = future + line_chars[i]
  }
  return future
}
