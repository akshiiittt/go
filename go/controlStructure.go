package main

import "fmt"

func controlStructure() {
	// fmt.Println("Hello")

	// if x := 1; x > 7 && x <= 12 {
	// 	fmt.Println("x is greater than 7")
	// } else if x < 3 && x >= 1 {
	// 	fmt.Println("x is less than 3")
	// } else {
	// 	fmt.Println("enter number bw 1-12")
	// }

	// char := 'a'
	// switch char {
	// case 'a', 'e', 'i', 'o', 'u':
	// 	fmt.Println("It's a vowel.")
	// case 'y':
	// 	fmt.Println("It's sometimes a vowel, sometimes a consonant.")
	// default:
	// 	fmt.Println("It's a consonant.")
	// }

	// switch day := "monday"; day {
	// case "monday":
	// 	fmt.Println("time to work!")
	// 	fallthrough
	// case "friday":
	// 	fmt.Println("let's party")
	// 	fallthrough
	// default:
	// 	fmt.Println("browse memes")

	// }

	// x := 3
	// switch {
	// case x > 6:
	// 	fmt.Println("x is greater than 6")
	// default:
	// 	fmt.Println("x is less than 6")
	// }

	// infinte loop
	// for{
	// 	fmt.Println("sjjdsnjk")
	// }

	//basic for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// while loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		if i == 6 {
			goto skip
		}

		i++
	}

skip:
	fmt.Println("*******************************")

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// for k, v := range arr {
	// 	fmt.Println(k, v)
	// }

	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		fmt.Println("Breaking the loop at i =", i)
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		fmt.Println("Skipping iteration at i =", i)
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// message := "Hello, Go!"
	// for index, char := range message {
	// 	fmt.Printf("Index: %d, Character: %c\n", index, char)
	// }

}
