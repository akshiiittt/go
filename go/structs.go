package main

import (
	"fmt"
	// "unsafe"
)

type Model struct {
	Name      string
	value     int
	completed bool
}

func (m Model) printName() {
	fmt.Println(m.completed)
}

func (m *Model) updateName(name string) {
	m.Name = name
}

func structs() {

	var car Model
	car.Name = "abc"
	car.value = 9
	car.completed = true
	fmt.Println(car)
	fmt.Println(car.Name)
	car.printName()
	car.updateName("aaa")
	fmt.Println(car)

	// var car2 = Model{"Abc", 99, true}
	// fmt.Println(car2)

	// // var car3 = Model{}
	// var car3 = Model{name: "abc"}
	// fmt.Println(car3)

	// var lang = struct {
	// 	Name string
	// }{"Golang"}

	// fmt.Println(lang)

	// var car5 = &car3
	// car5.completed = true
	// fmt.Println(car5)
	// fmt.Println(car3)
	// fmt.Println(car3 == *car5)

	// // exported with Capital letters and small letters with in the same file

	// var s struct{}
	// fmt.Println(unsafe.Sizeof(s))

	// type Outer struct {
	// 	one string
	// 	two string
	// }

	// type Inner struct {
	// 	Outer
	// 	three string
	// }

	// var value1 Inner
	// value1.one = "one"
	// value1.two = "two"
	// value1.three = "three"
	// fmt.Println(value1)
	// fmt.Println(value1.Outer)

}
