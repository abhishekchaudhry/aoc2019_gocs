package main

import (
"fmt"
"io/ioutil"
"strconv"
"strings"
)

type cordinates struct {
	x, y int
}

func main() {

	input, _ := ioutil.ReadFile("input/day3_dummy.txt")
	wires := strings.Fields(string(input))
	var wirePoints []map[cordinates]int
	var stepCounts []map[cordinates]int
	var commonPoints []cordinates

	for _, wire := range wires {
		path := make(map[cordinates]int)
		steps := make(map[cordinates]int)
		X := 0
		Y := 0
		step := 0
		moves := strings.Split(wire, ",")
		for _, move := range moves {
			direction := string(move[0])
			distance, _ := strconv.Atoi(move[1:])
			for i := 0; i < distance; i++ {
				switch direction {
				case "R":
					X++
				case "L":
					X--
				case "U":
					Y++
				case "D":
					Y--
				}
				path[cordinates{x: X, y: Y}]++
				step++
				_, ok := steps[cordinates{x: X, y: Y}]
				if !ok {
					steps[cordinates{x: X, y: Y}] = step
				}

			}
		}

		wirePoints = append(wirePoints, path)
		stepCounts = append(stepCounts, steps)
	}

	for k1 := range wirePoints[0] {
		if _, ok := wirePoints[1][k1]; ok {
			commonPoints = append(commonPoints, k1)
		}
	}

	var min int
	for _, v := range commonPoints {
		mDist := stepCounts[0][v] + stepCounts[1][v]
		if min == 0 || int(mDist) < min {
			min = int(mDist)
		}
	}

	fmt.Println(min)
}