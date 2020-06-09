package main

import (
	"fmt"
	"log"
	"math"
)

func getSizeByUnit(size int, unit string) string {
	var exp int
	switch unit {
	case "auto":
		return getAutoSize(size)
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
		log.Printf("Warning: unknown unit '%s' in getSizeByUnit(), defaulting to B\n", unit)
		unit = " B"
		exp = 1
	}

	return fmt.Sprintf("%.3f %s", float64(size)*math.Pow10(exp), unit)
}

func getAutoSize(size int) string {
	unitNo := 0
	var unit string
	sizeFloat := float64(size)
	for sizeFloat > 1000 {
		sizeFloat = sizeFloat / 1000
		unitNo++
	}

	switch unitNo {
	case 0:
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
		log.Fatal("Error in getAutoSize(): this shouldn't have happened because 64bit integers can't reach sizes larger than ~10^18")
	}

	return fmt.Sprintf("%.3f %s", sizeFloat, unit)
}
