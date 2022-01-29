package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
	float32 | float64 
}

func InsertionSort[number Number](input []number) []number {
	var n = len(input)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j = j - 1
		}
	}
	return input
}

func main() {
	list := []int32{4,3,1,5,}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}

	sorted := InsertionSort(list)
	fmt.Println(sorted)

	sortedFloats := InsertionSort(list2)
	fmt.Println(sortedFloats)
}