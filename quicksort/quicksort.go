package main

import "fmt"

type Number interface {
	int16 | int32 | int64 | float32 | float64 
}

func Partition[N Number](input []N, low, high int) ([]N, int) {
	pivot := input[high]
	i := low
	for j := low; j < high; j++ {
		if input[j] < pivot {
			input[i], input[j] = input[j], input[i]
			i++
		}
	} 
	input[i], input[high] = input[high], input[i]
	return input, i
}

func QuickSort[N Number](input []N, low, high int) []N {
	if low < high {
		input, partition := Partition(input, low, high)
		input = QuickSort(input, low, partition-1)
		input = QuickSort(input, partition+1, high)
	}
	return input
}

func main() {
	listInt := []int32{3,5,1,2,6,7,4,2,}
	listFloat := []float32{3.3,.5,1.5,22.3,64.2,7.1,1.4,2,}
	sorted := QuickSort(listInt, 0, len(listInt)-1)
	fmt.Println(sorted)

	sortedFloats := QuickSort(listFloat, 0, len(listFloat)-1)
	fmt.Println(sortedFloats)

}