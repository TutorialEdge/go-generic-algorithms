package main

import "fmt"


type Number interface {
	int16 | int32 | int64 | float32 | float64 
}

func BubbleSortGeneric[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
		  	if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
		  	}
		}
	}
	return input
}


func main() {
	fmt.Println("Go Generics Tutorial")
	list := []int32{4,3,1,5,}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}
	sorted := BubbleSortGeneric(list)
	fmt.Println(sorted)

	sortedFloats := BubbleSortGeneric(list2)
	fmt.Println(sortedFloats)

}
