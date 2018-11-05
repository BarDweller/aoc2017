package main

import (
	"fmt"
	"strings"
	"strconv"
)

func jmp(program []int, pc int) int{
	steps := program[pc]
	
	offset := 1
	if steps>=3 {
		offset = -1
	}
	program[pc] += offset
	
	return steps
}

func runProg(program []int) int {	
	steps,pc:=0,0
	for steps,pc = 0,0 ; (pc>=0 && pc<len(program)) ; pc,steps=pc+jmp(program,pc),steps+1 {
		//fmt.Println(steps," At ",pc," doing jmp(",program[pc],")")
		//fmt.Println(program)
	}
	return steps
}

func loadProg(input string) []int {
	lines := strings.Split(input, "\n")
	program := []int{}
	for _, line := range lines {
		if line == "" { continue }	
		if val,err := strconv.Atoi(strings.TrimSpace(line)); err==nil {
			program = append(program,val)
		}else{
			panic(line)
		}
	}
	return program
}

func doDay5(input string){
	program := loadProg(input)
	steps := runProg(program);
	fmt.Println("done",steps,program)	
}


func main() {
	input := `
0
3
0
1
-3
`
	doDay5(input)
}
