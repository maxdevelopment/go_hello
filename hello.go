package main

import (
	"fmt"
	"math/rand"
)

func main() {

	array := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("created array: %v \narray length: %v\n\n", array, len(array))

	slice := array[1:5]
	fmt.Printf("created slice: %v \nslice length: %v\nslice capacity: %v\n\n", slice, len(slice), cap(slice))

	slice[1] = 33
	fmt.Printf("changed slice element: %v \nslice length: %v\nslice capacity: %v\n\n", slice, len(slice), cap(slice))

	slice = append(slice, 55)
	fmt.Printf("added slice element: %v \nslice length: %v\nslice capacity: %v\n\n", slice, len(slice), cap(slice))

	arrayMod := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("array for modifications: %v\n", arrayMod)

	for i, j := 0, len(arrayMod)-1; i < j; i, j = i+1, j-1 {
		arrayMod[i], arrayMod[j] = arrayMod[j], arrayMod[i]
	}
	fmt.Printf("reversed: %v\n", arrayMod)

	for i := range arrayMod {
		j := rand.Intn(i + 1)
		arrayMod[i], arrayMod[j] = arrayMod[j], arrayMod[i]
	}
	fmt.Printf("shuffled: %v\n", arrayMod)
}
