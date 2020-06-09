package main

import (
	"fmt"
	"log"
	"math"
)

// GetSizeByUnit converts bytes to any other units ranging from B to EB.
// Utilizes GetSizeAuto for the "auto" unit. Returns a string composed of
// a size (float64 with 3 decimal places) and the corresponding unit.
func GetSizeByUnit(size int, unit string) string {
	var exp int
	switch unit {
	case "auto":
		return GetSizeAuto(size)
	case "B":
		exp = 1
	case "kB":
		exp = -3
	case "MB":
		exp = -6
	case "GB":
		exp = -9
	case "TB":
		exp = -12
	case "PB":
		exp = -15
	case "EB":
		exp = -18
	// no need for more units because 10^18 approaches the limits of 64bit integers
	default:
		log.Printf("Warning: unknown unit '%s' in GetSizeByUnit(), defaulting to B\n", unit)
		unit = " B"
		exp = 1
	}

	return fmt.Sprintf("%.3f %s", float64(size)*math.Pow10(exp), unit)
}

// GetSizeAuto converts bytes to a unit such that the resulting number will be
// between 1 (inclusive) and 1000 (exclusive). Returns a string composed of a size
// (float64 with 3 decimal places) and the corresponding unit.
func GetSizeAuto(size int) string {
	unitNo := 0
	var unit string
	sizeFloat := float64(size)
	for sizeFloat > 1000 {
		sizeFloat = sizeFloat / 1000
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
		log.Fatal("Error in GetSizeAuto(): this shouldn't have happened because 64bit integers can't reach sizes larger than ~10^18")
	}

	return fmt.Sprintf("%.3f %s", sizeFloat, unit)
}
