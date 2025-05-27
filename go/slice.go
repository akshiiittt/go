package main

import (
	"fmt"
)

func slice() {
	fmt.Println("hellooooo")

	// slices in the are teh three things in the struct and they point ot kind of the same thing in if mutiple thing are pointed to it so they point ot the same thing

	// type slice struct {
	// 	array unsafe.Pointer
	// 	len   int
	// 	cap   int
	// }

	arr := []int{}
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(arr == nil)

	arr1 := make([]int, 2, 3)
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1, len(arr1), cap(arr1))

	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	fmt.Println(nilSlice == nil)

	var emptySlice = []int{}
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	fmt.Println(emptySlice == nil)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Println(slice2, len(slice2), cap(slice2))
	// slice2 := slice1[1:4]
	// fmt.Println(slice2, len(slice2), cap(slice2))
	// slice2 := slice1[2:4]
	// fmt.Println(slice2, len(slice2), cap(slice2))
	// len = high - low
	// cap = cap(number) - low

	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num[1:2], len(num[1:2]), cap(num[1:2]), "...")
	fmt.Println(num[:2], len(num[:2]), cap(num[:2]))
	fmt.Println(num[1:], len(num[1:]), cap(num[1:]))

	var ans = evenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(ans)

	originalArray := []int{1, 2, 3}
	deepCopy := make([]int, len(originalArray))
	copy(deepCopy, originalArray)
	deepCopy[0] = 99
	fmt.Println(deepCopy)
	fmt.Println(originalArray)

	concatSlice := append(originalArray, deepCopy...)
	fmt.Println(concatSlice)

	oldSlice := make([]int, 2, 3)
	oldSlice = append(oldSlice, 1)
	newSlice := append(oldSlice, 3) // now this will make the copy beacuse it is reallocating the to the new array
	newSlice[1] = 9
	// newSlice[1] = 9 this will change in both as they are pointing towards the same thing
	fmt.Println(oldSlice, len(oldSlice), cap(oldSlice))
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

}

func evenNumbers(numbers []int) []int {
	var evenArr []int
	for _, v := range numbers {
		if v%2 == 0 {
			evenArr = append(evenArr, v)
		}
	}
	return evenArr
}
