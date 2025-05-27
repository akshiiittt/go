package main

import (
	"fmt"
	// "time"
)

func gemini() {
	i := 9
	j := 9.9
	fmt.Println(i)
	fmt.Println(j)
	// fmt.Println(i + j)
	fmt.Println(float64(i) + j)

	// const currentTime = time.Now()

	const MyInt = 100 // MyInt is an untyped constant
	var a int = MyInt
	var b int32 = MyInt
	var c float64 = MyInt
	fmt.Println(a, b, c) // Output: 100 100 100

	const (
		// iota is 0 for Monday, then increments
		Monday    = iota // 0
		Tuesday          // 1 (implicitly Tuesday = iota)
		Wednesday        // 2
		Thursday         // 3
		Friday           // 4
		Saturday         // 5
		Sunday           // 6
	)

	const (
		// iota resets to 0 for a new const block
		Red   = iota // 0
		Blue         // 1
		Green        // 2
	)

	const (
		_      = iota             // Blank identifier '_' to skip the first iota value (0)
		KB int = 1 << (10 * iota) // 1 << (10 * 1) = 1024
		MB int = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
		GB int = 1 << (10 * iota) // 1 << (10 * 3)
	)

	fmt.Println("Monday:", Monday)   // 0
	fmt.Println("Tuesday:", Tuesday) // 1
	fmt.Println("Sunday:", Sunday)   // 6
	fmt.Println("Red:", Red)         // 0
	fmt.Println("MB:", MB)           // 1048576

	sum := 0
	for sum < 10 {
		fmt.Println(sum)
		sum++
	}

	numbers := []int{10, 20, 30, 40, 50}
	for index, values := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, values)
	}

	text := "hello"
	for i, r := range text {
		fmt.Println(i, r)
	}

	arr := [7]int{}
	for i, r := range arr {
		fmt.Println(i, r)
	}

	grade := "B"

	fmt.Println("Grade evaluation:")
	switch grade {
	case "A":
		fmt.Println("Excellent!")
		fallthrough
	case "B":
		fmt.Println("Good!")
		// Explicitly fall through to the next case
	case "C":
		fmt.Println("Passable.")
		fallthrough // Explicitly fall through
	case "D":
		fmt.Println("Needs improvement.")
	default:
		fmt.Println("Invalid grade.")
	}
	fmt.Println("Evaluation finished.")

	char := 'a' // 'a' is a rune in Go
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("It's a vowel.")
	case 'y':
		fmt.Println("It's sometimes a vowel, sometimes a consonant.")
	default:
		fmt.Println("It's a consonant.")
	}

	cities := [...]string{"London", "Paris", "Tokyo"}
	fmt.Println(cities)

	arr1 := [...]int{1, 2, 3}
	fmt.Println(arr1)

	var numbers1 [8]int
	fmt.Println(numbers1)

	var names [3]string
	fmt.Println(names)

	colors := [3]string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("Color at index %d: %s\n", i, color)
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	// mixed := []interface{}{1, "hello", true} // This works only because type is interface{}
	// fmt.Println(mixed...)

	var slice1 = make([]int, 3, 3)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(slice1)

}

// Output:
// Grade evaluation:
// Good!
// Passable.
// Evaluation finished.

func main() {
	gemini()
}

// ----------------
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     fmt.Println("Correct way to use loop variable in goroutine:")
//     nums := []int{1, 2, 3}
//     for _, n := range nums {
//         // Correct: Create a new variable 'numCopy' for each iteration
//         numCopy := n
//         go func() {
//             time.Sleep(10 * time.Millisecond) // Simulate some work
//             fmt.Println(numCopy)
//         }()
//     }
//     time.Sleep(100 * time.Millisecond) // Wait for goroutines to finish

//     fmt.Println("\nIncorrect way (shows final value):")
//     for _, n := range nums {
//         // Incorrect: 'n' is reused, goroutines might see the final value (3)
//         go func() {
//             time.Sleep(10 * time.Millisecond)
//             fmt.Println(n) // This will likely print 3, 3, 3 or some mix
//         }()
//     }
//     time.Sleep(100 * time.Millisecond)
// }
