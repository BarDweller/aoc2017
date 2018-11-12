package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Registers map[string]int
type Prog []Line
type Line struct {
	reg        string
	op         string
	val        int
	cmd        string
	compareReg string
	compareOp  string
	compareVal int
}

func findMax(regs Registers) int {
	max := 0
	for _, r := range regs {
		if r > max {
			max = r
		}
	}
	return max
}

func runLine(regs Registers, line Line) {
	test := false
	switch line.compareOp {
	case ">":
		test = regs[line.compareReg] > line.compareVal
	case "<":
		test = regs[line.compareReg] < line.compareVal
	case ">=":
		test = regs[line.compareReg] >= line.compareVal
	case "<=":
		test = regs[line.compareReg] <= line.compareVal
	case "==":
		test = regs[line.compareReg] == line.compareVal
	case "!=":
		test = regs[line.compareReg] != line.compareVal
	}

	if test {
		switch line.op {
		case "inc":
			regs[line.reg] += line.val
		case "dec":
			regs[line.reg] -= line.val
		}
	}
}

func runProg(prog Prog) {
	regs := Registers{}
	for _, line := range prog {
		runLine(regs, line)
	}
	max := findMax(regs)
	fmt.Println("Max: ", max)
}

func asInt(txt string) int {
	if val, err := strconv.Atoi(txt); err == nil {
		return val
	} else {
		panic("bad input, expected number, got : " + txt)
	}
}

func parseLine(linestr string) Line {
	parts := strings.Split(linestr, " ")
	line := Line{}

	line.reg = parts[0]
	line.op = parts[1]
	line.val = asInt(parts[2])
	line.cmd = parts[3]
	line.compareReg = parts[4]
	line.compareOp = parts[5]
	line.compareVal = asInt(parts[6])

	return line
}

func loadData(input string) Prog {
	lines := []Line{}
	for _, l := range strings.Split(input, "\n") {
		t := strings.TrimSpace(l)
		if t == "" {
			continue
		}
		line := parseLine(t)
		lines = append(lines, line)
	}
	return lines
}

func main() {
	input := testdata()
	prog := loadData(input)
	runProg(prog)
	fmt.Println("Done")
}

func testdata() string {
	return `
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10    
    `
}
