package main

import (
	"fmt"
	"math"
	"testing"
)

type getSizeByUnitMock struct {
	size int
	unit string
	want string
	ok   bool
}

type getSizeAutoMock struct {
	size int
	want string
	ok   bool
}

var validUnits = []string{
	"auto", "B", "kB", "MB", "GB", "TB", "PB", "EB"}

var invalidUnits = []string{
	"YB", "BYTE", "hello", "go"}

var sizesLtZero = []int{
	-1, -100, -999, -1000, -1001, -4500, -149386, -1647583927}

var sizeGtZeroData = []getSizeByUnitMock{
	{100, "auto", fmt.Sprintf("%.3f  B", 100*math.Pow10(0)), true},
	{100, "B", fmt.Sprintf("%.3f  B", 100*math.Pow10(0)), true},
	{100, "kB", fmt.Sprintf("%.3f kB", 100*math.Pow10(-3)), true},
	{100, "MB", fmt.Sprintf("%.3f MB", 100*math.Pow10(-6)), true},
	{100, "GB", fmt.Sprintf("%.3f GB", 100*math.Pow10(-9)), true},
	{100, "TB", fmt.Sprintf("%.3f TB", 100*math.Pow10(-12)), true},
	{100, "PB", fmt.Sprintf("%.3f PB", 100*math.Pow10(-15)), true},
	{100, "EB", fmt.Sprintf("%.3f EB", 100*math.Pow10(-18)), true},
	{100, "YB", "", false},
	{100, "BYTE", "", false},
	{100, "hello", "", false},
	{100, "go", "", false},

	{999, "auto", fmt.Sprintf("%.3f  B", 999*math.Pow10(0)), true},
	{999, "B", fmt.Sprintf("%.3f  B", 999*math.Pow10(0)), true},
	{999, "kB", fmt.Sprintf("%.3f kB", 999*math.Pow10(-3)), true},
	{999, "MB", fmt.Sprintf("%.3f MB", 999*math.Pow10(-6)), true},
	{999, "GB", fmt.Sprintf("%.3f GB", 999*math.Pow10(-9)), true},
	{999, "TB", fmt.Sprintf("%.3f TB", 999*math.Pow10(-12)), true},
	{999, "PB", fmt.Sprintf("%.3f PB", 999*math.Pow10(-15)), true},
	{999, "EB", fmt.Sprintf("%.3f EB", 999*math.Pow10(-18)), true},
	{999, "YB", "", false},
	{999, "BYTE", "", false},
	{999, "hello", "", false},
	{999, "go", "", false},

	{1000, "auto", fmt.Sprintf("%.3f kB", 1000*math.Pow10(-3)), true},
	{1000, "B", fmt.Sprintf("%.3f  B", 1000*math.Pow10(0)), true},
	{1000, "kB", fmt.Sprintf("%.3f kB", 1000*math.Pow10(-3)), true},
	{1000, "MB", fmt.Sprintf("%.3f MB", 1000*math.Pow10(-6)), true},
	{1000, "GB", fmt.Sprintf("%.3f GB", 1000*math.Pow10(-9)), true},
	{1000, "TB", fmt.Sprintf("%.3f TB", 1000*math.Pow10(-12)), true},
	{1000, "PB", fmt.Sprintf("%.3f PB", 1000*math.Pow10(-15)), true},
	{1000, "EB", fmt.Sprintf("%.3f EB", 1000*math.Pow10(-18)), true},
	{1000, "YB", "", false},
	{1000, "BYTE", "", false},
	{1000, "hello", "", false},
	{1000, "go", "", false},

	{1001, "auto", fmt.Sprintf("%.3f kB", 1001*math.Pow10(-3)), true},
	{1001, "B", fmt.Sprintf("%.3f  B", 1001*math.Pow10(0)), true},
	{1001, "kB", fmt.Sprintf("%.3f kB", 1001*math.Pow10(-3)), true},
	{1001, "MB", fmt.Sprintf("%.3f MB", 1001*math.Pow10(-6)), true},
	{1001, "GB", fmt.Sprintf("%.3f GB", 1001*math.Pow10(-9)), true},
	{1001, "TB", fmt.Sprintf("%.3f TB", 1001*math.Pow10(-12)), true},
	{1001, "PB", fmt.Sprintf("%.3f PB", 1001*math.Pow10(-15)), true},
	{1001, "EB", fmt.Sprintf("%.3f EB", 1001*math.Pow10(-18)), true},
	{1001, "YB", "", false},
	{1001, "BYTE", "", false},
	{1001, "hello", "", false},
	{1001, "go", "", false},

	{51839, "auto", fmt.Sprintf("%.3f kB", 51839*math.Pow10(-3)), true},
	{51839, "B", fmt.Sprintf("%.3f  B", 51839*math.Pow10(0)), true},
	{51839, "kB", fmt.Sprintf("%.3f kB", 51839*math.Pow10(-3)), true},
	{51839, "MB", fmt.Sprintf("%.3f MB", 51839*math.Pow10(-6)), true},
	{51839, "GB", fmt.Sprintf("%.3f GB", 51839*math.Pow10(-9)), true},
	{51839, "TB", fmt.Sprintf("%.3f TB", 51839*math.Pow10(-12)), true},
	{51839, "PB", fmt.Sprintf("%.3f PB", 51839*math.Pow10(-15)), true},
	{51839, "EB", fmt.Sprintf("%.3f EB", 51839*math.Pow10(-18)), true},
	{51839, "YB", "", false},
	{51839, "BYTE", "", false},
	{51839, "hello", "", false},
	{51839, "go", "", false},

	{189792058, "auto", fmt.Sprintf("%.3f MB", 189792058*math.Pow10(-6)), true},
	{189792058, "B", fmt.Sprintf("%.3f  B", 189792058*math.Pow10(0)), true},
	{189792058, "kB", fmt.Sprintf("%.3f kB", 189792058*math.Pow10(-3)), true},
	{189792058, "MB", fmt.Sprintf("%.3f MB", 189792058*math.Pow10(-6)), true},
	{189792058, "GB", fmt.Sprintf("%.3f GB", 189792058*math.Pow10(-9)), true},
	{189792058, "TB", fmt.Sprintf("%.3f TB", 189792058*math.Pow10(-12)), true},
	{189792058, "PB", fmt.Sprintf("%.3f PB", 189792058*math.Pow10(-15)), true},
	{189792058, "EB", fmt.Sprintf("%.3f EB", 189792058*math.Pow10(-18)), true},
	{189792058, "YB", "", false},
	{189792058, "BYTE", "", false},
	{189792058, "hello", "", false},
	{189792058, "go", "", false},

	{597540590326, "auto", fmt.Sprintf("%.3f GB", 597540590326*math.Pow10(-9)), true},
	{597540590326, "B", fmt.Sprintf("%.3f  B", 597540590326*math.Pow10(0)), true},
	{597540590326, "kB", fmt.Sprintf("%.3f kB", 597540590326*math.Pow10(-3)), true},
	{597540590326, "MB", fmt.Sprintf("%.3f MB", 597540590326*math.Pow10(-6)), true},
	{597540590326, "GB", fmt.Sprintf("%.3f GB", 597540590326*math.Pow10(-9)), true},
	{597540590326, "TB", fmt.Sprintf("%.3f TB", 597540590326*math.Pow10(-12)), true},
	{597540590326, "PB", fmt.Sprintf("%.3f PB", 597540590326*math.Pow10(-15)), true},
	{597540590326, "EB", fmt.Sprintf("%.3f EB", 597540590326*math.Pow10(-18)), true},
	{597540590326, "YB", "", false},
	{597540590326, "BYTE", "", false},
	{597540590326, "hello", "", false},
	{597540590326, "go", "", false},

	{45555555555555, "auto", fmt.Sprintf("%.3f TB", 45555555555555*math.Pow10(-12)), true},
	{45555555555555, "B", fmt.Sprintf("%.3f  B", 45555555555555*math.Pow10(0)), true},
	{45555555555555, "kB", fmt.Sprintf("%.3f kB", 45555555555555*math.Pow10(-3)), true},
	{45555555555555, "MB", fmt.Sprintf("%.3f MB", 45555555555555*math.Pow10(-6)), true},
	{45555555555555, "GB", fmt.Sprintf("%.3f GB", 45555555555555*math.Pow10(-9)), true},
	{45555555555555, "TB", fmt.Sprintf("%.3f TB", 45555555555555*math.Pow10(-12)), true},
	{45555555555555, "PB", fmt.Sprintf("%.3f PB", 45555555555555*math.Pow10(-15)), true},
	{45555555555555, "EB", fmt.Sprintf("%.3f EB", 45555555555555*math.Pow10(-18)), true},
	{45555555555555, "YB", "", false},
	{45555555555555, "BYTE", "", false},
	{45555555555555, "hello", "", false},
	{45555555555555, "go", "", false},

	{22222222222222222, "auto", fmt.Sprintf("%.3f PB", 22222222222222222*math.Pow10(-15)), true},
	{22222222222222222, "B", fmt.Sprintf("%.3f  B", 22222222222222222*math.Pow10(0)), true},
	{22222222222222222, "kB", fmt.Sprintf("%.3f kB", 22222222222222222*math.Pow10(-3)), true},
	{22222222222222222, "MB", fmt.Sprintf("%.3f MB", 22222222222222222*math.Pow10(-6)), true},
	{22222222222222222, "GB", fmt.Sprintf("%.3f GB", 22222222222222222*math.Pow10(-9)), true},
	{22222222222222222, "TB", fmt.Sprintf("%.3f TB", 22222222222222222*math.Pow10(-12)), true},
	{22222222222222222, "PB", fmt.Sprintf("%.3f PB", 22222222222222222*math.Pow10(-15)), true},
	{22222222222222222, "EB", fmt.Sprintf("%.3f EB", 22222222222222222*math.Pow10(-18)), true},
	{22222222222222222, "YB", "", false},
	{22222222222222222, "BYTE", "", false},
	{22222222222222222, "hello", "", false},
	{22222222222222222, "go", "", false},
}

// size < 0
func TestGetSizeByUnitSizeLtZero(t *testing.T) {
	for _, size := range sizesLtZero {
		for _, unit := range append(validUnits, invalidUnits...) {
			got, err := GetSizeByUnit(size, unit)
			if err == nil {
				t.Errorf(
					"GetSizeByUnit(): %d -> %s, got %q, expected non-nil error",
					size, unit, got,
				)
				t.Fail()
			}
		}
	}
}

// size == 0
func TestGetSizeByUnitSizeEqZero(t *testing.T) {
	for _, unit := range validUnits {
		got, err := GetSizeByUnit(0, unit)
		if unit == "auto" || unit == "B" {
			unit = " B"
		}
		want := fmt.Sprintf("0.000 %s", unit)
		if got != want || err != nil {
			t.Errorf("GetSizeByUnit(): 0 -> %s, got %q, expected %q with nil error",
				unit, got, want)
			t.Fail()
		}
	}
}

// size > 0
func TestGetSizeByUnitSizeGtZero(t *testing.T) {
	for _, sample := range sizeGtZeroData {
		got, err := GetSizeByUnit(sample.size, sample.unit)
		var ok bool
		var nilStr string
		if err == nil {
			ok = true
			nilStr = "nil"
		} else {
			ok = false
			nilStr = "non-nil"
		}
		want := sample.want
		if got != want || ok != sample.ok {
			t.Errorf("GetSizeByUnit(): %d -> %s, got %q, expected %q with %s error",
				sample.size, sample.unit, got, want, nilStr)
			t.Fail()
		}
	}
}
