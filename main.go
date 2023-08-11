package main

import "fmt"

func from_string(str string) (any, []token) {
	tokens := Lex(str)
	fmt.Println("tokens : ", tokens)

	val, _ := Parse(tokens, true)

	return val, tokens
}

func main() {

	str := "{\"key\":{\"key2\":\"val2\", \"key3\": \"val3\"}, \"key4\":[2,4]}"

	val, _ := from_string(str)

	fmt.Println("val: ", val)
}
