package main

import (
	"fmt"
	"strings"
)

const DIGITS string = "0123456789-"

type JsonValueType string

const (
	Number  JsonValueType = "0123456789-"
	String                = "\""
	Boolean               = "true|false"
	Array                 = "[]"
)

type JsonValue struct {
	typeof JsonValueType
	value  any
}

type ValuePair struct {
	key   string
	value JsonValue
}

// type ObjectPair struct {
// 	key    string
// 	object Object
// }

type Object struct {
	pairs []ValuePair
}

type ArrayPair struct {
	values []JsonValue
}

var lines []string
var chars []string
var char string
var idx int
var col int

var lineNum int

var tokens []Object

func parse() {
	lines = strings.Split(content, EOL)
	for l, line := range lines {
		lineNum = l
		if len(line) > 0 {
			chars = strings.Split(line, "")
			idx = 0
			char = chars[idx]
			col = 1
			collectObject()
		}
	}
}

func collectObject() {
	var currentPair ValuePair
	var currentKey string
	var currentValue JsonValue
	var currentObject Object
	for char != "" {
		printCurrentChar()
		if char == " " || char == "\t" || char == EOL {
			advance()
		} else if char == "{" {
			// start new object
			advance()
		} else if char == "\"" {
			// start pair
			currentKey = collectString()
			fmt.Println("key", currentKey)
		} else if char == ":" {
			// start value
			currentValue = collectValue()
			fmt.Println("value", currentValue)
			if len(currentKey) > 0 {
				currentPair = ValuePair{
					key:   currentKey,
					value: currentValue,
				}
				currentObject.pairs = append(currentObject.pairs, currentPair)
			}
		} else if char == "," || char == "}" {
			// end pair or object
			advance()
		} else {
			panic(fmt.Sprintf("Unknown char: %s (%d:%d)", getCurrentChar(), lineNum+1, col))
		}
	}
	tokens = append(tokens, currentObject)
}

func collectValue() (PairValue JsonValue) {
	fmt.Println("collecting value")
	advance()
	var typeof JsonValueType
	var value any
	for char != "," && char != "]" && char != "" {
		if char == " " || char == "\t" || char == EOL {
			advance()
		} else if char == "\"" {
			fmt.Println("collecting string value")
			value = collectString()
			typeof = String
		} else if strings.ContainsAny(DIGITS, char) {
			fmt.Println("collecting number value")
			value = collectNumber()
			typeof = Number
		} else if char == "t" || char == "f" {
			fmt.Println("collecting boolean value")
			if char == "t" {
				value = "true"
				advanceTimes(4)
			} else if char == "f" {
				value = "false"
				advanceTimes(5)
			}
			typeof = Boolean
		} else if char == "[" {
			fmt.Println("collecting array value")
			value = collectArray()
			typeof = Array
			advance()
		} else {
			panic(fmt.Sprintf("Unknown character in value: %s (%d:%d)", getCurrentChar(), lineNum+1, col))
		}
	}
	fmt.Println("END VALUE", typeof, value)
	PairValue = JsonValue{
		typeof: typeof,
		value:  value,
	}
	return
}

func collectArray() (array ArrayPair) {
	for char != "]" {
		fmt.Println("collecting array value char")
		array.values = append(array.values, collectValue())
		fmt.Println(array)
	}
	return
}

func collectNumber() (number string) {
	fmt.Println("collecting number")
	for strings.ContainsAny(DIGITS, char) {
		number += char
		advance()
	}
	return
}

func collectString() (str string) {
	advance()
	fmt.Println("collecting string")
	for char != "\"" && prevChar(1) != "\\" {
		printCurrentChar()
		str += char
		advance()
	}
	fmt.Println("collected string")
	advance()
	return
}

func advance() {
	col++
	idx++
	if len(chars) > idx {
		char = chars[idx]
	} else {
		char = ""
	}
}

func advanceTimes(times int) {
	for i := 0; i < times; i++ {
		advance()
	}
}

func prevChar(mov int) (prevChar string) {
	var prevCharIdx int = idx
	if len(chars) < (idx - mov) {
		prevChar = " "
		for prevChar != " " {
			prevCharIdx -= mov
			prevChar = chars[prevCharIdx]
		}
	}
	return
}

func nextChar(mov int) (nextChar string) {
	var nextCharIdx int = idx
	if len(chars) < (idx + mov) {
		nextChar = " "
		for nextChar != " " {
			nextCharIdx += mov
			nextChar = chars[nextCharIdx]
		}
	}
	return
}

func printCurrentChar() {
	var currentChar string
	switch char {
	case "\t":
		currentChar = "TAB"
		break
	case " ":
		currentChar = "SPACE"
		break
	case EOL:
		currentChar = "EOL"
		break
	case "":
		currentChar = "(empty)"
	default:
		currentChar = char
	}
	fmt.Println(currentChar)
}

func getCurrentChar() (currentChar string) {
	switch char {
	case "\t":
		currentChar = "\\t"
		break
	case " ":
		currentChar = "(space)"
		break
	case EOL:
		currentChar = "\\n"
		break
	case "":
		currentChar = "(empty)"
	default:
		currentChar = char
	}
	return
}
