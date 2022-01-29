package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 |
	float32 | float64 
}

func MergeSort[number Number](input []number) []number {
	if len(input) < 2 {
		return input
	}
	first := MergeSort(input[:len(input)/2])
	second := MergeSort(input[len(input)/2:])
	return Merge(first, second)
}

func Merge[number Number](a, b []number) []number {
	final := []number{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}


func main() {
	list := []int32{4,3,1,5,}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}

	sorted := MergeSort(list)
	fmt.Println(sorted)

	sortedFloats := MergeSort(list2)
	fmt.Println(sortedFloats)
}