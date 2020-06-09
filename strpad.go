package main

import (
	"fmt"
	"log"
)

// Strlpad pads a copy of the input string on the left side with spaces to
// make its length equal to  padLen. If the string's initial length is greater
// or equal to padLen, the same string is returned.
func Strlpad(str string, padLen int) string {
	result, err := strpad(str, padLen, "left")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	return result
}

// Strrpad pads a copy of the input string on the left side with spaces to
// make its length equal to  padLen. If the string's initial length is greater
// or equal to padLen, the same string is returned.
func Strrpad(str string, padLen int) string {
	result, err := strpad(str, padLen, "right")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	return result
}

func strpad(str string, padLen int, padDirection string) (string, error) {
	if padLen <= len(str) {
		return str, nil
	}

	// create byte slice of spaces
	whitespace := make([]byte, padLen-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

	// append spaces to the input string
	bytes := []byte(str)
	switch padDirection {
	case "left":
		bytes = append(whitespace, bytes...)
	case "right":
		bytes = append(bytes, whitespace...)
	default:
		return str, fmt.Errorf("strpad(): invalid pad direction '%s'", padDirection)
	}

	return string(bytes), nil
}
