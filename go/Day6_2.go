
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	orbits := make(map[string]string)
	neworbits := map[string]int{}

	file, err := os.Open("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ")")
		orbits[line[1]] = line[0]
	}


	for key,ok := orbits["YOU"];ok;{
		fmt.Println(key)
		if ok {
			neworbits[key] = len(neworbits)
			key = orbits[key]
				if key == ""{
					break
				}
		} else {
			break
		}
	}

	distance := 0
	for orbit, ok := orbits["SAN"]; ok; orbit, ok = orbits[orbit] {
		if _, ok := neworbits[orbit]; ok {
			fmt.Println(distance + neworbits[orbit])
			break
		}
		distance++
	}

}