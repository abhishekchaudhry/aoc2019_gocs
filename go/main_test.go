package main

import "testing"

func Test_checkFuel(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"12",args{12},2},
		{"14",args{14},2},
		{"1969",args{1969},654},
		{"100756",args{100756},33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkFuel(tt.args.m); got != tt.want {
				t.Errorf("checkFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkTotalFuel(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"12",args{12},2},
		{"1969",args{1969},966},
		{"100756",args{100756},50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTotalFuel(tt.args.m); got != tt.want {
				t.Errorf("checkTotalFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}