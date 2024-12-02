package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := os.ReadFile("numbers.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	// init slice for left numbers and list for right numbers
	var list1, list2 []int
	counts := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var n1, n2 int
		// format string line into integers and append into slice
		fmt.Sscanf(s, "%d %d", &n1, &n2)
		list1, list2 = append(list1, n1), append(list2, n2)
		counts[n2]++
	}

	// sort slice of integers lowest to greatest
	slices.Sort(list1)
	slices.Sort(list2)

	// init count
	part1, part2 := 0, 0
	for i := range list1 {
		// find abs difference between the two integers and add to part 1 total
		part1 += abs(list2[i] - list1[i])
		part2 += list1[i] * counts[list1[i]]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
