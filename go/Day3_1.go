package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coordinates struct{
	position complex64
	elapsed int
}

type navPoints map[complex64]int

func initNavpoints() navPoints{
	return make(navPoints)
}

func (set navPoints) intersect(secondSet navPoints) navPoints {
	intersections := initNavpoints()
	for pos := range set{
		if _,ok := secondSet[pos]; ok{
			intersections.add(pos,set[pos]+secondSet[pos])
		}
	}
	return intersections
}

func (set navPoints) toSlice() []coordinates {
	retval := []coordinates{}
	for pos := range set {
		retval = append(retval, coordinates{position: pos, elapsed: set[pos]})
	}
	return retval
}

func (set navPoints) add(points complex64, elapsed int){
	if _,ok := set[points]; !ok {
		set[points] = elapsed
	}

}

func manhattanDist(points coordinates) int {
	dist := math.Abs(float64(real(points.position))) + math.Abs(float64(imag(points.position)))
	return int(dist)
}

type ManDist []coordinates

func (points ManDist) Len() int           { return len(points) }
func (points ManDist) Less(i, j int) bool { return manhattanDist(points[i]) < manhattanDist(points[j]) }
func (points ManDist) Swap(i, j int)      { points[i], points[j] = points[j], points[i] }

type ElapsedDist []coordinates

func (points ElapsedDist) Len() int           { return len(points) }
func (points ElapsedDist) Less(i, j int) bool { return points[i].elapsed < points[j].elapsed }
func (points ElapsedDist) Swap(i, j int)      { points[i], points[j] = points[j], points[i] }

func readInput(path string)([]string, []string){
	file,err :=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	List_A := strings.Split(scanner.Text(), ",")

	//fmt.Print("the first list is :", List_A)

	scanner.Scan()

	List_B := strings.Split(scanner.Text(),",")

	//fmt.Print("the second list is :", List_B)

	return List_A, List_B
}

func createPath(navpath []string) navPoints{
	var position complex64 = 0
	elapsed := 0
	path := initNavpoints()

	for _,value := range navpath{
		var nav complex64
		switch first_char := string(value[0]); {
		case first_char == "L":
			nav = -1
		case first_char == "R":
			nav = 1
		case first_char == "U":
			nav = 1i
		case first_char == "D":
			nav = -1i
		default:
			fmt.Println("Invalid data in path for direction")
			fmt.Println(value)
		}

		dist,err := strconv.Atoi(value[1:])
		if err!=nil{
			panic(err)
		}

		for i:=0;i<dist;i++ {
			position += nav
			elapsed++
			path.add(position,elapsed)
			}

		}
	return path
}

func findIntersection(pathA navPoints,pathB navPoints) coordinates{
	intersections := pathA.intersect(pathB).toSlice()

	sort.Sort(ManDist(intersections))

	return intersections[0]
}

func findIntersectionElapsed(pathA navPoints,pathB navPoints) coordinates{
	intersections := pathA.intersect(pathB).toSlice()

	sort.Sort(ElapsedDist(intersections))

	return intersections[0]
}

func popPath(listA []string, listB []string) coordinates{
	pathA :=createPath(listA)
	pathB :=createPath(listB)

	//fmt.Println(pathA,pathB)
	return findIntersection(pathA,pathB)
}

func popPathElapsed(listA []string, listB []string) coordinates{
	pathA :=createPath(listA)
	pathB :=createPath(listB)

	//fmt.Println(pathA,pathB)
	return findIntersectionElapsed(pathA,pathB)
}

func main() {
	listA, listB := readInput("input/day3.txt")
	fmt.Println("Manhattan distance is")
	fmt.Println(manhattanDist(popPath(listA,listB)))

	fmt.Println("Part 2 elapsed distance is")
	fmt.Println(popPathElapsed(listA,listB))

}
