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
	var input="1212"
	//input="1221"
	//input="123425"
	//input="123123"
	//input="12131415"
	
	var offset=len(input)/2
	input+=input
	
	var total uint64	
	for i:=0 ; i<len(input)/2; i++ {
	  if input[i] == input[i+offset] {
	    total += charToInt(rune(input[i]))
	  }
	}
	
	fmt.Println("Total",total)	
}
