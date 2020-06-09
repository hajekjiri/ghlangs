package main

func strlpad(str string, pad int) string {
	if pad <= len(str) {
		return string(str)
	}

	whitespace := make([]byte, pad-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

	bytes := []byte(str)
	bytes = append(whitespace, bytes...)
	return string(bytes)
}

func strrpad(str string, pad int) string {
	if pad <= len(str) {
		return string(str)
	}

	whitespace := make([]byte, pad-len(str))
	for i := range whitespace {
		whitespace[i] = ' '
	}

	bytes := []byte(str)
	bytes = append(bytes, whitespace...)
	return string(bytes)
}
