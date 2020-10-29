package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkFuel(m int) int{
	return (m / 3) -2
}

func main() {

	file,err := os.Open("input/day1.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()


	total :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		var mass int
		_, err := fmt.Sscanf(line,"%d", &mass)
		if err != nil {
			panic(err)
		}

		fuel := checkFuel(mass)
		for fuel >= 0 {
			total += fuel
			fuel = checkFuel(fuel)
		}


		//fuel += (mass/3) -2
	}

	fmt.Println(" fuel needed for second case",total)
}
