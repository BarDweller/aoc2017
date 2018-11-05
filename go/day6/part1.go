package main

import (
	"fmt"
	"strings"
	"strconv"
)

func loadBanks(input string)[]int{
	banks := []int{}
	for _,line := range strings.Split(input,"\n") {
		for _,countStr := range strings.Split(line,"	") {
			if count,err := strconv.Atoi(countStr); err==nil {
				banks = append(banks,count);
			}		
		}
	}
	return banks
}

func findCandidate(banks []int)int{
	max:=0
	selected:=0
	for idx,count := range banks {
		if count>max {
			selected=idx
			max=count
		}
	}
	return selected
}

func redistributeBanks(banks []int)int{
	seen:=map[string]bool{}
	cycles:=0
	for ;; {				
		key := strings.Join(strings.Fields(fmt.Sprint(banks)),":")
		if _, ok := seen[key]; ok {
			return cycles;
		}else{
			seen[key]=true
		}
		
		fmt.Println(banks)
		
		idx:=findCandidate(banks)				
		count:=banks[idx]
		banks[idx]=0
		
		fmt.Println("Selected Bank ",idx," with count ",count)
	
		for c,i:=count,idx+1; c>0; c,i=c-1,i+1 {
			addIdx := i%len(banks)
			banks[addIdx] = banks[addIdx]+1
		}
		
		cycles++		
	}
}

func doDay6(input string) {
	banks := loadBanks(input)
	cycles := redistributeBanks(banks)
	fmt.Println(cycles)
}

func main() {
	input:=`
14	0	15	12	11	11	3	5	1	6	8	4	9	1	8	4
`
	doDay6(input)
}
