package main

import (
	"testing"
)

// pad < len(str)
func TestPadLtStrLen(t *testing.T) {
	str := "hello"
	padLen := len("hello") - 2
	if got, want := Strlpad(str, padLen), "hello"; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "hello"; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad == len(str)
func TestPadEqStrLen(t *testing.T) {
	str := "hello"
	padLen := len("hello")
	if got, want := Strlpad(str, padLen), "hello"; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "hello"; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad > len(str)
func TestPadGtStrLen(t *testing.T) {
	str := "hello"
	padLen := len("hello") + 3
	if got, want := Strlpad(str, padLen), "   hello"; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "hello   "; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad < 0 && len(str) > 0
func TestPadLtZeroStrLenGtZero(t *testing.T) {
	str := "hello"
	padLen := -3
	if got, want := Strlpad(str, padLen), "hello"; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "hello"; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad < 0 && len(str) == 0
func TestPadLtZeroStrLenEqZero(t *testing.T) {
	str := ""
	padLen := -3
	if got, want := Strlpad(str, padLen), ""; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), ""; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad == 0 && len(str) > 0
func TestPadEqZeroStrLenGtZero(t *testing.T) {
	str := "hello"
	padLen := 0
	if got, want := Strlpad(str, padLen), "hello"; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "hello"; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad == 0 && len(str) == 0
func TestPadEqZeroStrLenEqZero(t *testing.T) {
	str := ""
	padLen := 0
	if got, want := Strlpad(str, padLen), ""; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), ""; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}

// pad > 0 && len(str) == 0
func TestPadGtZeroStrLenEqZero(t *testing.T) {
	str := ""
	padLen := 4
	if got, want := Strlpad(str, padLen), "    "; got != want {
		t.Errorf("Strlpad(): got %q, expected %q", got, want)
		t.Fail()
	}

	if got, want := Strrpad(str, padLen), "    "; got != want {
		t.Errorf("Strrpad(): got %q, expected %q", got, want)
		t.Fail()
	}
}
