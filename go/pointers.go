package main

import "fmt"

func pointers() {
	var x *int
	fmt.Println(x) // <nil>

	a := 9
	b := &a
	fmt.Println(b) // 0x1400000e0d8

	var c *int = &a
	fmt.Println(c)  // 0x1400011e010
	fmt.Println(*c) // 9

	*c = 20
	fmt.Println(a)  // 20
	fmt.Println(*b) // 20
	fmt.Println(*c) // 20

	p := new(int)
	*p = 100
	fmt.Println(p)
	fmt.Println(*p)

	p1 := &p
	fmt.Println(p1)
	fmt.Println(**p1)

	// valueOne := 6
	// addOne(&valueOne)
	// addOne(&valueOne)
	// fmt.Println(valueOne) // 7

}

// func addOne(value *int) {
// 	*value++
// }
