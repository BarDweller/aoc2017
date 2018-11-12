package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getLengths(input string) []int {
	result := []int{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for _, valstr := range strings.Split(line, ",") {
			trimmed := strings.TrimSpace(valstr)
			if val, err := strconv.Atoi(trimmed); err == nil {
				result = append(result, val)
			} else {
				panic("Bad input " + valstr)
			}
		}
	}
	return result
}

func getList(count int) []int {
	result := []int{}
	for i := 0; i < count; i++ {
		result = append(result, i)
	}
	return result
}

func hash(list []int, lengths []int) {
	skipSize := 0
	currentIndex := 0
	listLen := len(list)
	for _, length := range lengths {
		if length > listLen {
			continue
		}
		for i := 0; i < length/2; i++ {
			sourceIdx := (currentIndex + i) % listLen
			destIdx := (currentIndex + length - i - 1) % listLen
			tmp := list[destIdx]
			list[destIdx] = list[sourceIdx]
			list[sourceIdx] = tmp
		}
		currentIndex += (length + skipSize) % listLen
		skipSize++
	}
}

func main() {
	lengths := getLengths(testdata())
	list := getList(5)
	fmt.Println(lengths)
	fmt.Println(list)

	hash(list, lengths)
	fmt.Println(list)

	fmt.Println(list[0] * list[1])
}

func testdata() string {
	return `
3, 4, 1, 5
	`
}
