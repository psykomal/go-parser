package main

import (
	"reflect"
	"testing"
)

type lexTest struct {
	input    string
	expected []token
}

var lexTests = []lexTest{
	{
		"{}",
		[]token{
			{SyntaxKind, string(JSON_LEFTBRACE)},
			{SyntaxKind, string(JSON_RIGHTBRACE)},
		},
	},
	{
		"{\"key\": \"value\"}",
		[]token{
			{SyntaxKind, string(JSON_LEFTBRACE)},
			{StringKind, "key"},
			{SyntaxKind, string(JSON_COLON)},
			{StringKind, "value"},
			{SyntaxKind, string(JSON_RIGHTBRACE)},
		},
	},
	{
		"{\"key\":{\"key2\":\"val2\", \"key3\": \"val3\"}, \"key4\":[2,4]}",
		[]token{
			{SyntaxKind, string(JSON_LEFTBRACE)},
			{StringKind, "key"},
			{SyntaxKind, string(JSON_COLON)},
			{SyntaxKind, string(JSON_LEFTBRACE)},
			{StringKind, "key2"},
			{SyntaxKind, string(JSON_COLON)},
			{StringKind, "val2"},
			{SyntaxKind, string(JSON_COMMA)},
			{StringKind, "key3"},
			{SyntaxKind, string(JSON_COLON)},
			{StringKind, "val3"},
			{SyntaxKind, string(JSON_RIGHTBRACE)},
			{SyntaxKind, string(JSON_COMMA)},
			{StringKind, "key4"},
			{SyntaxKind, string(JSON_COLON)},
			{SyntaxKind, string(JSON_LEFTBRACKET)},
			{NumberKind, "2"},
			{SyntaxKind, string(JSON_COMMA)},
			{NumberKind, "4"},
			{SyntaxKind, string(JSON_RIGHTBRACKET)},
			{SyntaxKind, string(JSON_RIGHTBRACE)},
		},
	},
}

func TestLex(t *testing.T) {

	for _, test := range lexTests {
		if output := Lex(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Lex(%s) returned %v, expected %v", test.input, output, test.expected)
		}
	}
}
