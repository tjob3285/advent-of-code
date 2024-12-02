package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// input consists or many reports, one report per line
// each report is a list of numbers called levels separated by spaces
/*
	7 6 4 2 1
	1 2 7 8 9
	9 7 6 2 1
	1 3 2 4 5
	8 6 4 4 1
	1 3 6 7 9
*/
// This data consists of 6 reports (rows) with each report having x levels (cols)
// report is SAFE if all increasing or decreasing & two adjacent levels differ by at least 1 and at most 3
/*
   7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
   1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
   9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
   1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
   8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
   1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*/

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	safeReports := 0
	safeReports2 := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var report []int
		json.Unmarshal([]byte("["+strings.ReplaceAll(s, " ", ",")+"]"), &report)

		val := checkReport(report)
		if val {
			safeReports++
		}
		for i := range report {
			if checkReport(slices.Delete(slices.Clone(report), i, i+1)) {
				safeReports2++
				break
			}
		}
	}

	fmt.Println(safeReports)
	fmt.Println(safeReports2)
}

// check if report is all increase/decrease and adjacent differ between 1-3
func checkReport(r []int) bool {
	for i := 1; i < len(r); i++ {
		// v = 7 - 6 (1), 1*6-7 = -1 so decreasing list, check difference between 1-3/-1-(-3)
		if v := r[i] - r[i-1]; v*(r[1]-r[0]) <= 0 || v < -3 || v > 3 {
			return false
		}
	}

	return true
}
