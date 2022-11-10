package main

import (
	"fmt"

	mySlice "github.com/JuliaAbdullaeva/testtasks/slice"
)

func main() {
	a := []int{1, 5, 3, 4, 2, 5, 56, 7, 9, 8, 10, 2}
	b := []int{2, 14, 4, 6, 56, 2}
	result := mySlice.IntersectionSlices(a, b)
	fmt.Println(result)
}
