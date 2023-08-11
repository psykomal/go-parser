package main

import (
	"fmt"
	"testing"
)

type parseTest struct {
	input    string
	expected map[string]interface{}
}

var parseTests = []parseTest{
	{
		"{}",
		map[string]interface{}{},
	},
	{
		"{\"key\": \"value\"}",
		map[string]interface{}{
			"key": "value",
		},
	},
	{
		"{\"key\":{\"key2\":\"val2\", \"key3\": \"val3\"}, \"key4\":[2,4]}",
		map[string]interface{}{
			"key": map[string]interface{}{
				"key2": "val2",
				"key3": "val3",
			},
			"key4": []interface{}{2, 4},
		},
	},
}

func TestParse(t *testing.T) {

	for _, test := range parseTests {
		if output, _ := from_string(test.input); fmt.Sprint(output) != fmt.Sprint(test.expected) {
			t.Errorf("Lex(%s) returned %v, expected %v", test.input, output, test.expected)
		}
	}
}
