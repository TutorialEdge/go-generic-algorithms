package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
	float32 | float64 
}

func SelectionSort[number Number](input []number) []number {
	var n = len(input)
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i; j < n; j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	} 
	return input
}

func main() {
	list := []int32{4,3,1,5,}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}


	sorted := SelectionSort(list)
	fmt.Println(sorted)
	sortedFloats := SelectionSort(list2)
	fmt.Println(sortedFloats)

}