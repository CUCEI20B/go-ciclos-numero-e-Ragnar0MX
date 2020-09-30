package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	result := e(num)
	fmt.Printf("%.3f", result)
}

func e(num float64) float64 {
	var eu, i float64
	eu = 0
	for i = 0; i <= num; i++ {
		ft := fact(i)
		acomulado := eu + (1 / ft)
		eu = acomulado
	}
	return eu
}

func fact(num float64) float64 {
	var facto float64
	facto = 1
	for i := num; i > 0; i-- {
		facto = facto * i
	}
	return facto
}
