package main

import (
	"fmt"
	"math"
)

// GetSizeByUnit converts bytes to any other units ranging from B to EB.
// Utilizes GetSizeAuto for the "auto" unit. Returns a string composed of
// a size (float64 with 3 decimal places) and the corresponding unit.
func GetSizeByUnit(size int, unit string) (string, error) {
	if size < 0 {
		return "", fmt.Errorf("GetSizeByUnit(): size cannot be less than 0")
	}

	var exp float64
	switch unit {
	case "auto":
		result, err := getSizeAuto(size)
		if err != nil {
			return "", err
		}
		return result, nil
	case "B":
		exp = 0
	case "kB":
		exp = -10
	case "MB":
		exp = -20
	case "GB":
		exp = -30
	case "TB":
		exp = -40
	case "PB":
		exp = -50
	case "EB":
		exp = -60
	// no need for more units because 10^18 approaches the limits of 64bit integers
	default:
		return "", fmt.Errorf("GetSizeByUnit(): unknown unit %q", unit)
	}

	if unit == "B" {
		unit = " B"
	}

	return fmt.Sprintf("%.3f %s", float64(size)*math.Pow(2, exp), unit), nil
}

// getSizeAuto converts bytes to a unit such that the resulting number will be
// between 1 (inclusive) and 1024 (exclusive) as long as the input size is
// greater than zero. Returns a string composed of a size (float64 with 3
// decimal places) and the corresponding unit. Input of size 0 will return
// "0  B".
func getSizeAuto(size int) (string, error) {
	unitNo := 0
	var unit string
	sizeFloat := float64(size)
	for sizeFloat >= 1024 {
		sizeFloat = sizeFloat / 1024
		unitNo++
	}

	switch unitNo {
	case 0:
		// prepend a space to 'B' for prettier output
		unit = " B"
	case 1:
		unit = "kB"
	case 2:
		unit = "MB"
	case 3:
		unit = "GB"
	case 4:
		unit = "TB"
	case 5:
		unit = "PB"
	case 6:
		unit = "EB"
		// no need for more units because 10^18 approaches the limits of 64bit integers
	default:
		return "", fmt.Errorf("getSizeAuto(): size is larger than the 64bit integer limit (?)")
	}

	return fmt.Sprintf("%.3f %s", sizeFloat, unit), nil
}
