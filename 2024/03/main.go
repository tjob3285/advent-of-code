package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// multiply some numbers like mul(X,Y) where X and Y are 1-3 digit numbers
// many invalid chars that should be ignored mul(4*, mul(6,9!, ?(12,34) or mul( 2 , 4 ) do nothing
// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
// only mul(2,4), mul(5,5), mul(11,8), mul(8,5) work -> 8+25+88+40 = 161

func main() {

	// read file contents
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	re := regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)
	fmt.Println(mul(string(input)))
	fmt.Println(mul(strings.Join(re.FindAllString(string(input), -1), "")))
}

func mul(s string) (r int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		n1, _ := strconv.Atoi(m[1])
		n2, _ := strconv.Atoi(m[2])
		r += n1 * n2
	}
	return r
}
