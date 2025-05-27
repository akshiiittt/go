package main

import "fmt"

// "fmt"
// "maps"
// "time"

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(1)
	// fmt.Println(true)
	// fmt.Println(9.9)

	// var name string = "Hello"
	// var lang string = "golang"
	// var isTrue bool = true
	// var num int = 9
	// var num1 = 9
	// shortHand := "values"
	// var values int
	// values = 999

	// const pi = 3.14

	// const (
	// 	host = "localhost"
	// 	port = 3000
	// )

	// var (
	// 	v1 = "one"
	// 	v2 = "two"
	// )

	// fmt.Println(name)
	// fmt.Println(lang)
	// fmt.Println(isTrue)
	// fmt.Println(num)
	// fmt.Println(num1)
	// fmt.Println(shortHand)
	// fmt.Println(values)
	// fmt.Println(pi)
	// fmt.Println(host)
	// fmt.Println(port)
	// fmt.Println(v1)
	// fmt.Println(v2)

	// var i = 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i2 := 10
	// for i2 > 0 {
	// 	fmt.Println(i2)
	// 	i2--
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("This is the for loop skjbjhbsdj")
	// }

	// for i := range 7 {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// value := 10
	// if value >= 6 {
	// 	fmt.Println("value is greater than 6")
	// } else if value < 6 {
	// 	fmt.Println("value is less than 6")
	// }

	// value = 20
	// fmt.Println(value)

	// score := 85

	// if score >= 90 {
	// 	fmt.Println("Grade: A")
	// } else if score >= 80 && score < 90 {
	// 	fmt.Println("Grade: B")
	// } else if score >= 70 && score < 80 {
	// 	fmt.Println("Grade: C")
	// } else if score >= 60 && score < 70 {
	// 	fmt.Println("Grade: D")
	// } else {
	// 	fmt.Println("Grade: F")
	// }

	// // var oneValue = "1"
	// //  oneValue = 1
	// //  fmt.Println(oneValue)

	// var val string = "admin"
	// var hasPermissions bool = true

	// if val == "admin" && hasPermissions {
	// 	fmt.Println("these are logical operators")
	// }

	// if val := 9; val > 8 {
	// 	fmt.Println("declaration inside if")
	// }

	// x := 3
	// switch x {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("This is the default value")
	// }

	// // fmt.Println(time.Now())

	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("workday")
	// }

	// // var k = "types"
	// // fmt.Println(strings.Split(k, ""))

	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("its int")
	// 	case string:
	// 		fmt.Println("its string")
	// 	case bool:
	// 		fmt.Println("its string")
	// 	default:
	// 		fmt.Println("its other variable", t)
	// 	}
	// }

	// whoAmI(9)

	// var nums [6]int
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4
	// nums[4] = 5
	// nums[5] = 6

	// fmt.Println(len(nums))
	// fmt.Println(nums)
	// fmt.Println(nums[0])

	// var stringArr [3]string
	// stringArr[0] = "h"
	// stringArr[1] = "i"
	// stringArr[2] = "j"
	// fmt.Println(stringArr)

	// numsArray := [3]int{1, 2, 3}
	// fmt.Println(numsArray)

	// twoDArray := [3][3]int{{6, 7, 1}, {8, 9, 1}, {9, 9, 9}}
	// fmt.Println(twoDArray)

	// var slices []int
	// fmt.Println(slices)
	// fmt.Println(slices == nil)
	// fmt.Println(len(slices))

	// var slices1 = make([]int, 2, 3)
	// // var slices1 = []int{}
	// slices1[0] = 1
	// slices1[1] = 2
	// // fmt.Println(slices1)
	// // fmt.Println(cap(slices1))
	// slices1 = append(slices1, 7)
	// slices1 = append(slices1, 8)
	// slices1 = append(slices1, 9)
	// slices1 = append(slices1, 10)
	// slices1 = append(slices1, 11)
	// slices1 = append(slices1, 12)
	// fmt.Println(slices1)
	// // fmt.Println(cap(slices1))

	// var sumArray = []int{1, 2, 3, 4, 5}
	// fmt.Println(sumArray[0:3])

	// arr := [3]int{10, 20, 30}
	// for i, val := range arr {
	// 	fmt.Println(i, val)
	// }

	// // var numsArray1 = [10]int{1,2,3,7:100, 8}
	// // fmt.Println(numsArray1)

	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["name1"] = "js"
	// fmt.Println(m)
	// fmt.Println(m["name"])

	// delete(m, "name1")
	// fmt.Println(m)

	// m2 := map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println(m2)

	// k, ok := m2["one"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("value there")
	// } else {
	// 	fmt.Println("value not thre")
	// }

	// m3 := map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println(maps.Equal(m2, m3))

	// string1 := "abcdefghi"
	// for k, v := range string1 {
	// 	fmt.Println(k, string(v))
	// }

	// add := func(a int, b int) int {
	// 	return a + b
	// }

	// result := add(1, 2)
	// fmt.Println(result)

	// getLanguages := func() (string, string, string) {
	// 	return "go", "js", "python"
	// }

	// l1, l2, l3 := getLanguages()
	// fmt.Println(l1, l2, l3)

	// numsArray1 := []int{1, 2, 3}
	// sumAns := sum(numsArray1...)
	// // sumAns := sum(1,2,3)
	// fmt.Println(sumAns)

	// increment := counter()
	// fmt.Println(increment())
	// fmt.Println(increment())

	// num := 1
	// changeNum(&num)
	// fmt.Println(num)

	order := newStrcut{
		ID:    1,
		name:  "abc",
		value: "one",
	}

	fmt.Println(order)
	order.ID = 9

	fmt.Println(order)
	fmt.Println(order.name)

	order.changeValue("hello")
	fmt.Println(order)

	a := make([]string, 3)
	a[0] = "a"
	a = append(a, "b")
	fmt.Println(a, cap(a))
}

type newStrcut struct {
	ID    int
	name  string
	value string
}

func (n *newStrcut) changeValue(value string) {
	n.value = value
}

// func sum(nums ...int) int {
// 	sum := 0
// 	for _, v := range nums {
// 		sum = sum + v
// 	}
// 	return sum
// }

// func counter() func() int {
// 	count := 0

// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func changeNum(num *int) {
// 	*num = 5
// 	fmt.Println(*num)
// }
