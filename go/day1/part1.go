package main

import (
	"fmt"
)

func charToInt(c rune) uint64 {
  var current uint64
  switch c {
	    case '0': current=0
	    case '1': current=1
	    case '2': current=2
	    case '3': current=3
	    case '4': current=4
	    case '5': current=5
	    case '6': current=6
	    case '7': current=7
	    case '8': current=8
	    case '9': current=9 
  }
  return current
}

func main() {
	var input = "1122"
	//input = "91212129"
	//input = "1111"
  
	var first= charToInt(rune(input[0:1][0]))
	var total uint64
	
	for _,nextchar := range input[1:] {
	  var asnum = charToInt(rune(nextchar))
	  if first == asnum {
	    total+=first
	  }
	  first = asnum
	}
	if first == charToInt(rune(input[0:1][0])) {
	  total+=first
	}
	
	fmt.Println("Total",total)	
}
