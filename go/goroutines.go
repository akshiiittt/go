package main

import (
	"fmt"
)

func speak(arg string) {
	fmt.Println(arg)
}
func goroutine() {
	go speak("Hellooooooosdjhbsjhbsbjhdsbjhdss")

}
