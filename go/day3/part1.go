package main

import (
	"fmt"
)

type Direction int
const (
  Down Direction=0
  Right Direction=1
  Left Direction=2
  Up Direction=3
)

func determineSideCount(targetval int) int{
  i:=1
  for ; (i*i)<=targetval; i+=2 {}
  return i;
}

func nextDir(last Direction) Direction{
  switch last {
    case Up: return Left
    case Left: return Down
    case Down: return Right
    case Right: return Up
  }
  panic("unknown direction")
}

func abs(val int) int {
  if val<0 { return -val}
  return val
}

func main() {
	targetval := 312051
		
	sideCount := determineSideCount(targetval)
	startVal := ((sideCount-2)*(sideCount-2))+1
	
	dir := Up;
	perSide:=sideCount-1
	x:=perSide/2
	y:=(perSide/2)
	count:=perSide
	for i:=startVal; i<=targetval; i++ {
	  switch dir {
	    case Up: y--
	    case Left: x--
	    case Down: y++
	    case Right: x++
	  }
	  count--;
	  if count==0 {
	    count=perSide
	    dir=nextDir(dir)
	  }	  
	}	
	fmt.Println("done",targetval,x,y,abs(x)+abs(y))
}
