package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

func testData() []string {
  data := []string {
"5 1 9 5",
"7 5 3",
"2 4 6 8"}	
  return data
}

func processLine(line string) uint64 {
  parts := strings.Split(line,"	")
  var max uint64
  var min uint64 = math.MaxUint64

  for _,val := range parts {
    if i,err := strconv.Atoi(val); err == nil {
	bigi := uint64(i)
	if bigi>max { max=bigi }
	if bigi<min { min=bigi }
    }else{
      panic("Bad input")
    }
    
  }

  return max-min
}

func main() {
	var lines = testData()
	
	var total uint64
	for i,_ :=range lines {
	  total+=processLine(lines[i])
	}
	fmt.Println(total)	
}
