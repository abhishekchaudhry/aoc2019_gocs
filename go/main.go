package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file,err := os.Open("input/day1.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	fuel := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		var mass int
		_, err := fmt.Sscanf(line,"%d", &mass)
		if err != nil {
			panic(err)
		}

		fuel += (mass/3) -2
	}

	fmt.Println(fuel)
}
