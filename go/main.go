package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkFuel(m int) int{
	return (m / 3) -2
}

func checkTotalFuel(m int) int{
	fuel:= (m / 3) -2
	if fuel > 0 {
		fuel += checkTotalFuel(fuel)
	} else {
		fuel =0
	}
	return fuel
}

func Day1_1() {
	file, err := os.Open("input/day1.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()


	total :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		var mass int
		_, err := fmt.Sscanf(line, "%d", &mass)
		if err != nil {
			panic(err)
		}

		total += checkFuel(mass)

		//fuel += (mass/3) -2
	}

	fmt.Printf("DAY 1_1 : Fuel needed : %d\n", total)
}
func Day1_2() {
	file, err := os.Open("input/day1.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()


	total :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		var mass int
		_, err := fmt.Sscanf(line, "%d", &mass)
		if err != nil {
			panic(err)
		}

		fuel:= checkTotalFuel(mass)
		total += fuel
	}

	fmt.Printf("DAY 1_2 : Fuel needed : %d\n", total)
}

func Day2_1() {

	bytes ,err := ioutil.ReadFile("input/day2.txt")
	if err != nil{
		panic(err)
	}

	var program []int
	input_string :=string(bytes)
	input_array := strings.Split(input_string,",")
	for _, contents := range input_array{
		temp,err := strconv.Atoi(contents)
		if err !=nil{
			panic(err)
		}
		program = append(program,temp)
		//fmt.Println("the index is ",program)
		//fmt.Println("and value is ",contents)
	}

	program[1] = 12
	program[2] = 2
	//fmt.Println("the value is",program[1])
	br:=false
	for pos :=0; pos < len(program); pos = pos + 4{
		switch program[pos]{
		case 1:
			program[program[pos+3]] = program[program[pos+1]] + program[program[pos+2]]
			//fmt.Println("the value is",program[pos+3])
		case 2:
			program[program[pos+3]] = program[program[pos+1]] * program[program[pos+2]]
			//fmt.Println("the value is",program[pos+3])
		case 99:
			br=true
		default:
			fmt.Printf("invalid option code at position 0 : %d\n", program[0])
		}
		if br == true{
			break
		}
	}
	fmt.Printf("DAY 2_1 : Value at position 0  : %d\n", program[0])

}

func Day2_2() {

	bytes ,err := ioutil.ReadFile("input/day2.txt")
	if err != nil{
		panic(err)
	}

	var program []int
	input_string :=string(bytes)
	input_array := strings.Split(input_string,",")
	for _, contents := range input_array{
		temp,err := strconv.Atoi(contents)
		if err !=nil{
			panic(err)
		}
		program = append(program,temp)
	}

	//program[1] = 12
	//program[2] = 2
	//fmt.Println("the value is",program[1])
	br:=false
	result := 0
	for n :=0; n < 100; n++ {
		for v := 0; v <100; v++ {
			temp := append([]int(nil), program...)
			temp[1] = n
			temp[2] = v
			br = false
			for pos :=0; pos < len(temp); pos += 4{
				switch temp[pos]{
				case 1:
					temp[temp[pos+3]] = temp[temp[pos+1]] + temp[temp[pos+2]]
					//fmt.Println("the value is",temp[pos+3])
				case 2:
					temp[temp[pos+3]] = temp[temp[pos+1]] * temp[temp[pos+2]]
					//fmt.Println("the value is",temp[pos+3])
				case 99:
					br=true
				//default:
					//continue
					//fmt.Printf("invalid option code at position 0 : %d\n", temp[0])
				}
				if temp[0] == 19690720{
					result = (100 * n) + v
					fmt.Printf("DAY 2_2 : Value at position 0  : %d\n", result)
				}
				if br == true{
					break
				}

			}
		}
	}

	//fmt.Printf("DAY 2_2 : Value at position 0  : %d\n", result)

}




func main() {

	Day1_1()
	Day1_2()
	Day2_1()
	Day2_2()

}


