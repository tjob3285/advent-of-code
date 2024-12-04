package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

// find XMAS, can be horizontal, vertical, diagonal, written backwards, or even overlapping other words
// need to examine 4x4 matrices?
func main() {
	// read file contents
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	grid := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	total := 0
	total2 := 0
	for p := range grid {
		total += strings.Count(strings.Join(adjacent(p, 4, grid), " "), "XMAS")
		total2 += strings.Count("AMAMASASAMAMAS", strings.Join(adjacent(p, 2, grid)[:4], ""))
	}

	fmt.Println(total)
	fmt.Println(total2)
}

func adjacent(p image.Point, i int, grid map[image.Point]rune) []string {
	d := []image.Point{
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}

	words := make([]string, len(d))
	for l, s := range d {
		for n := range i {
			words[l] += string(grid[p.Add(s.Mul(n))])
		}
	}

	return words
}
