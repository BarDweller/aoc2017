package main

import (
	"fmt"
	"strings"
	"strconv"
)

func runProg(program []int) int {	
	steps,pc:=0,0
	for steps,pc = 0,0 ; (pc>=0 && pc<len(program)) ; pc,steps=pc+program[pc]-1,steps+1 {
		program[pc]+=1;
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
	fmt.Println("done",steps)	
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
