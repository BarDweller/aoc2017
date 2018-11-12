package main

import (
	"fmt"
	"strings"
)

func getLengths(input string) []int {
	result := []int{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for _, valrune := range line {
			result = append(result, int(valrune))
		}
	}
	result = append(result, 17)
	result = append(result, 31)
	result = append(result, 73)
	result = append(result, 47)
	result = append(result, 23)
	return result
}

func getList(count int) []int {
	result := []int{}
	for i := 0; i < count; i++ {
		result = append(result, i)
	}
	return result
}

func hash(list []int, lengths []int, rounds int) {
	listLen := len(list)
	currentIndex := 0
	skipSize := 0
	for round := 0; round < rounds; round++ {
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
}

func calcDenseHash(list []int) {
	blockSize := 16
	result := make([]int, len(list)/blockSize)
	for i := 0; i < len(result); i++ {
		val := 0
		for j := 0; j < blockSize; j++ {
			val ^= list[(i*blockSize)+j]
		}
		result[i] = val
	}
	for _, val := range result {
		fmt.Printf("%02x", val)
	}
}

func main() {
	lengths := getLengths(testdata())
	list := getList(256)

	hash(list, lengths, 64)

	calcDenseHash(list)
}

func testdata() string {
	return `
1,2,4
	`
}
