package main

import (
	"fmt"
	"strings"
)

func processLine(input string) (int, int, int) {
	skip := false
	inGarbage := false
	groupCount := 0
	groupScore := 0
	depth := 0
	garbageCount := 0
	for _, r := range input {
		if skip {
			skip = false
		} else {
			if inGarbage {
				garbageCount++
				switch r {
				case '!':
					garbageCount--
					skip = true
				case '>':
					garbageCount--
					inGarbage = false
				}
			} else {
				switch r {
				case '!':
					skip = true
				case '<':
					inGarbage = true
				case '{':
					depth++
				case '}':
					groupCount++
					groupScore += depth
					depth--
				}
			}
		}
	}

	return groupCount, groupScore, garbageCount
}

func processData(input string) {
	for idx, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		groups, groupScore, garbageCount := processLine(line)
		fmt.Println(line)
		fmt.Println("Input ", idx, " had ", groups, " groups with score ", groupScore, " with garbageCount ", garbageCount)
	}
}

func main() {
	input := testdata()
	processData(input)
}

func testdata() string {
	return `
<>
<random characters>
<<<<>
<{!>}>
<!!>
<!!!>>
<{o"i!a,<{i<a>
	`
}
