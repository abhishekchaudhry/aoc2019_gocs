package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

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

		fmt.Printf("the value at position 0 : %d\n", program[0])
	}








}
