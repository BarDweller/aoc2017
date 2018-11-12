package main

import (
	"fmt"
	"strings"
)

func processLine(input string) (int, int) {
	skip := false
	inGarbage := false
	groupCount := 0
	groupScore := 0
	depth := 0
	for _, r := range input {
		if skip {
			skip = false
		} else {
			if inGarbage {
				switch r {
				case '!':
					skip = true
				case '>':
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

	return groupCount, groupScore
}

func processData(input string) {
	for idx, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		groups, groupScore := processLine(line)
		fmt.Println("Input ", idx, " had ", groups, " groups with score ", groupScore)
	}
}

func main() {
	input := testdata()
	processData(input)
}

func testdata() string {
	return `
{}
{{{}}}
{{},{}}
{{{},{},{{}}}}
{<a>,<a>,<a>,<a>}
{{<ab>},{<ab>},{<ab>},{<ab>}}
{{<!!>},{<!!>},{<!!>},{<!!>}}
{{<a!>},{<a!>},{<a!>},{<ab>}}
`}
