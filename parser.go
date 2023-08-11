package main

import (
	"fmt"
)

func parse_array(tokens []token) (interface{}, []token) {

	json_arr := make([]interface{}, 0)

	t := tokens[0]
	if t.val == string(JSON_RIGHTBRACKET) {
		return json_arr, tokens[1:]
	}

	for len(tokens) > 0 {
		json, newTokens := Parse(tokens, false)
		json_arr = append(json_arr, json)

		tokens = newTokens
		t = tokens[0]
		if t.val == string(JSON_RIGHTBRACKET) {
			return json_arr, tokens[1:]
		} else if t.val != string(JSON_COMMA) {
			panic(fmt.Sprintf("Expected comma after pair in map, got %v", t))
		} else {
			tokens = tokens[1:]
		}
	}

	panic("Expected end of array bracket")
}

func parse_map(tokens []token) (interface{}, []token) {
	json_map := make(map[string]interface{})

	t := tokens[0]
	if t.val == string(JSON_RIGHTBRACE) {
		return json_map, tokens[1:]
	}

	for len(tokens) > 0 {
		json_key := tokens[0]
		if json_key.typ == StringKind {
			tokens = tokens[1:]
		} else {
			panic("Expected string type for key")
		}

		json_val, newTokens := Parse(tokens[1:], false)

		json_map[json_key.val] = json_val

		tokens = newTokens

		t = tokens[0]
		if t.val == string(JSON_RIGHTBRACE) {
			return json_map, tokens[1:]
		} else if t.val != string(JSON_COMMA) {
			panic(fmt.Sprintf("Expected comma after pair in map, got %v", t))
		}

		tokens = tokens[1:]
	}

	panic("Expected end of map bracket")
}

func Parse(tokens []token, is_root bool) (interface{}, []token) {
	t := tokens[0]

	if is_root && !((t.val == string(JSON_LEFTBRACE)) || (t.val == string(JSON_LEFTBRACKET))) {
		panic("Root must be a map or array")
	}

	if t.val == string(JSON_LEFTBRACKET) {
		return parse_array(tokens[1:])
	} else if t.val == string(JSON_LEFTBRACE) {
		return parse_map(tokens[1:])
	} else {
		return t.val, tokens[1:]
	}
}
