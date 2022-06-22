package main

import (
	"fmt"
	"strings"
)

var c string
var pos int
var line_chars []string

const digits string = "0123456789"

func Lex(json string) {
	line_chars = strings.Split(json, "")
	c = line_chars[pos]
	for len(c) > 0 {
		if c != "\t" && c != " " {
			if c == "\n" {
				if len(line_chars) == (pos + 1) {
					tokens = append(tokens, "EOF")
				} else {
					tokens = append(tokens, "EOL")
				}
			} else if c == "{" {
				tokens = append(tokens, "LCURLY")
			} else if c == "}" {
				tokens = append(tokens, "RCURLY")
			} else if c == ":" {
				tokens = append(tokens, "COLON")
			} else if c == "\"" {
				tokens = append(tokens, "LQUOTE")
				advance()
				token_string()
			} else if c == "[" {
				tokens = append(tokens, "LBRACKET")
			} else if c == "]" {
				tokens = append(tokens, "RBRACKET")
			} else if c == "," {
				tokens = append(tokens, "COMMA")
			} else if c == "t" || c == "f" {
				token_bool()
			} else if strings.Contains(digits, c) {
				token_number()
			}
		}
		advance()
	}
}

func token_number() {
	var number string
	for strings.Contains(digits, c) {
		number = number + c
		advance()
	}
	number = fmt.Sprintf("INT:%s", number)
	tokens = append(tokens, number)
}

func token_string() {
	var str string
	for c != "\"" {
		str = str + c
		advance()
	}
	str = fmt.Sprintf("STRING:%s", str)
	tokens = append(tokens, str, "RQUOTE")
}

func token_bool() {
	var peek string = peek(4)
	if peek == "true" {
		tokens = append(tokens, "BOOL:true")
	} else if peek == "false" {
		tokens = append(tokens, "BOOL:false")
	}
}

func advance() {
	pos += 1
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
