package main

import (
	"fmt"
)

func main() {
	a := []int{1, 5, 3, 4, 2, 5, 56, 7, 9, 8, 10, 2}
	b := []int{2, 14, 4, 6, 56, 2}
	result := intersectionSlices(a, b)
	fmt.Println(result)
}

//На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
func intersectionSlices(a, b []int) (intersectSlice []int) {
	mapIntersect := make(map[int]int)
	for _, itemA := range a {
		for _, itemB := range b {
			if itemA == itemB {
				_, ok := mapIntersect[itemB]
				if !ok {
					mapIntersect[itemB] = itemB
					intersectSlice = append(intersectSlice, itemB)
				}
			}
		}
	}
	return
}
