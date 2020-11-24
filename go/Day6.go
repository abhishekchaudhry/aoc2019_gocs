package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){

	orbits := make(map[string]string)
	count :=0

	file,err := os.Open("input/day6_dummy.txt")
	if err!=nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(),")")
		orbits[line[1]]= line[0]
	}

	fmt.Println("total orbits input are :",orbits)

	for orbit := range orbits{
		key := orbit
		fmt.Println("the orbit is:",key)
		//if key = value then check if there are more values associated with that
		for{
			neworbit,ok := orbits[orbit]
			if ok {
				orbit = neworbit
				count++
			} else {
				break
			}
		}
	}
	fmt.Println("total count", count)
}