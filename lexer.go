package main

import "fmt"

var JSON_WHITESPACE = [5]byte{' ', '\t', '\b', '\n', '\r'}
var JSON_SYNTAX = [6]byte{JSON_COMMA, JSON_COLON, JSON_LEFTBRACKET, JSON_RIGHTBRACKET,
	JSON_LEFTBRACE, JSON_RIGHTBRACE}

const (
	FALSE_LEN = len("false")
	TRUE_LEN  = len("true")
	NULL_LEN  = len("null")
)

type tokenType uint

const (
	StringKind tokenType = iota
	NumberKind
	BoolKind
	NullKind
	SyntaxKind
)

type token struct {
	typ tokenType
	val string
}

func lex_string(str string, cur int) (token, int, bool) {
	json_string := ""

	if str[cur] == JSON_QUOTE {
		cur = cur + 1
	} else {
		return token{}, cur, false
	}

	for i := cur; i < len(str); i++ {
		c := str[i]
		if c == JSON_QUOTE {
			return token{StringKind, json_string}, cur + len(json_string) + 1, true
		}

		json_string += string(c)
	}

	panic("Expected closing quote")
}

func lex_number(str string, cur int) (token, int, bool) {
	json_number := ""

	number_chars := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '-', '+'}
	found := false

	for i := cur; i < len(str); i++ {
		c := str[i]
		found = false
		for j := 0; j < len(number_chars); j++ {
			num_char := number_chars[j]
			if c == num_char {
				json_number += string(c)
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	rest := cur + len(json_number)

	if len(json_number) == 0 {
		return token{}, rest, false
	}

	return token{NumberKind, json_number}, rest, true
}

func lex_bool(str string, cur int) (token, int, bool) {
	str_len := len(str) - cur

	if str_len >= TRUE_LEN && str[cur:cur+TRUE_LEN] == "true" {
		return token{BoolKind, "true"}, cur + TRUE_LEN, true
	} else if str_len >= FALSE_LEN && str[cur:cur+FALSE_LEN] == "false" {
		return token{BoolKind, "false"}, cur + FALSE_LEN, true
	} else {
		return token{}, cur, false
	}
}

func lex_null(str string, cur int) (token, int, bool) {
	str_len := len(str) - cur

	if str_len >= NULL_LEN && str[cur:cur+NULL_LEN] == "null" {
		return token{NullKind, "null"}, cur + NULL_LEN, true
	} else {
		return token{}, cur, false
	}
}

func Lex(str string) []token {
	var tokens []token
	var cursor = 0
	fmt.Printf("input: %v\n", str)

lex:
	for cursor < len(str) {
		// fmt.Println(cursor, len(str), string(str[cursor]))
		if json_string, newCursor, ok := lex_string(str, cursor); ok {
			cursor = newCursor
			tokens = append(tokens, json_string)
			// fmt.Println("string: ", json_string)
			continue lex
		}
		if json_number, newCursor, ok := lex_number(str, cursor); ok {
			cursor = newCursor
			tokens = append(tokens, json_number)
			// fmt.Println("number: ", json_number)
			continue lex
		}
		if json_bool, newCursor, ok := lex_bool(str, cursor); ok {
			cursor = newCursor
			tokens = append(tokens, json_bool)
			// fmt.Println("bool: ", json_bool)
			continue lex
		}
		if json_null, newCursor, ok := lex_null(str, cursor); ok {
			cursor = newCursor
			tokens = append(tokens, json_null)
			// fmt.Println("null: ", json_null)
			continue lex
		}

		c := str[cursor]
		for _, whitespace_char := range JSON_WHITESPACE {
			if c == whitespace_char {
				cursor = cursor + 1
				continue lex
			}
		}

		for _, syntax_char := range JSON_SYNTAX {
			if c == syntax_char {
				token := token{
					typ: SyntaxKind,
					val: string(c),
				}
				tokens = append(tokens, token)
				cursor = cursor + 1
				continue lex
			}
		}

		panic(fmt.Sprintf("unexpected character: %c", c))

	}

	return tokens
}
