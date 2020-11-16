package main

import (
	"fmt"
)

func IntToSlice(n int, sequence []int) []int{
	if n != 0 {
		i := n % 10
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func checkOrder (seq []int) bool{
	asc := true
	pair := false

	for key,value := range seq {
		if key == 0 { continue }
		lastDigit :=seq[key-1]
		if value < lastDigit { asc = false}
		if value == lastDigit { pair = true}
		}
		return asc && pair
}

func checkOrderpart2 (seq []int) bool {
	asc := true
	pair := false
	count := 0

	fmt.Println(seq[1])

	for i := 0; i < len(seq); i++ {
		if i == 0 {
			continue
		}
		lastdigit := seq[i-1]
		if seq[i] < lastdigit {
			asc = false
		}
		if seq[i] == lastdigit {
			pair = true
			count = count + 1
		}
	}

	/*	for key,value := range seq {
		if key == 0 { continue }
		lastDigit :=seq[key-1]
		if value < lastDigit { asc = false}
		if value == lastDigit { pair = true}
	}*/
	if count == 2 {
		return asc && pair
	} else {
		return false
	}
}
func main() {
	var input_seq []int
	count :=0
	for i := 124075; i < 580770; i++ {
		input_seq = append(input_seq, i)
	}

	for i:=0;i<len(input_seq);i++{
		var seq []int


		x := input_seq[i]
		seq = IntToSlice(x,seq)
		//fmt.Println(seq)
		//fmt.Println(seq[2])

		if checkOrder(seq){
			count++
		}

/*				if checkOrderpart2(seq){
				count++
			}*/

/*			for key:=0;key<6;key++{
				raising := 0
				pair := 0
				if key == 0 { continue }
					lastDigit :=seq[key-1]
					fmt.Println("last digit",lastDigit)
					fmt.Println("current value",seq[key])
				if seq[key]<lastDigit { raising = raising +1}
				if seq[key]==lastDigit { pair = pair +1}
					fmt.Println("value of raising",raising)
					fmt.Println("value of pair",pair)
			}*/
	}
	fmt.Println("total count is ",count)
}