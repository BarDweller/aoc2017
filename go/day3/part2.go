package main

import (
	"fmt"
	"strconv"
)

type Direction int

const (
	Down  Direction = 0
	Right Direction = 1
	Left  Direction = 2
	Up    Direction = 3
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up   "
	case Left:
		return "Left "
	case Down:
		return "Down "
	case Right:
		return "Right"
	}
	panic("unknown direction")
}

type Coord struct {
	x int
	y int
}

func nextDir(last Direction) Direction {
	switch last {
	case Up:
		return Left
	case Left:
		return Down
	case Down:
		return Right
	case Right:
		return Up
	}
	panic("unknown direction")
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func determineMaxExtent(grid map[Coord]int) int {
	max := 0
	for k, _ := range grid {		
		if abs(k.x) > max {
			max = abs(k.x)
		}else if abs(k.y) > max {
			max = abs(k.y)
		}		
	}
	return max
}

func ppGrid(grid map[Coord]int, target int) {
	sidecount := determineMaxExtent(grid)
	for y := -sidecount; y <= sidecount; y++ {
		row := ""
		for x := -sidecount; x <= sidecount; x++ {
			cell, ok := grid[Coord{x, y}]
			if ok {
				row += strconv.Itoa(cell) + "\t"
			} else {
				row += " \t"
			}

		}
		fmt.Println(row)
	}
}

func sumAdjacents(grid map[Coord]int, x, y int) int {
	var result int
	for xo := -1; xo < 2; xo++ {
		for yo := -1; yo < 2; yo++ {
			cell, ok := grid[Coord{x + xo, y + yo}]
			if ok {
				result += cell
			}
		}
	}
	return result
}

func main() {
	targetval := 312051	

	sideCount := 1

	dir := Right
	perSide := sideCount - 1
	x := 0
	y := 0
	count := perSide

	var grid = map[Coord]int{Coord{0, 0}: 1}
	value := 0

	for value < targetval {

		//fmt.Println("val ", i, " dir ", dir, " x ", x, " y ", y, " remaining ", count)

		value = sumAdjacents(grid, x, y)

		grid[Coord{x, y}] = value

		switch dir {
		case Up:
			y--
		case Left:
			x--
		case Down:
			y++
		case Right:
			x++
		}

		count--
		if count <= 0 {
			dir = nextDir(dir)
			switch dir {
			case Right:
				count = perSide + 1
			case Up:
				perSide += 2
				count = perSide - 1
			case Left:
				count = perSide
			case Down:
				count = perSide
			}
		}
	}
	fmt.Println("done: value: ",value)
	ppGrid(grid, targetval)
}
