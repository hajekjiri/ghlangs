package main

import (
	"fmt"
	"log"
)

// Strlpad pads a string on the left side by padLen
func Strlpad(str string, padLen int) string {
	result, err := strpad(str, padLen, "left")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	return result
}

// Strrpad pads a string on the left side by padLen
func Strrpad(str string, padLen int) string {
	result, err := strpad(str, padLen, "right")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

	return result
}

func strpad(str string, padLen int, padDirection string) (string, error) {
	if padLen <= len(str) {
		return string(str), nil
	}

	whitespace := make([]byte, padLen-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

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
