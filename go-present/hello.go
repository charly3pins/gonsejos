package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}

func highlighted() {
	fmt.Println("I am highlighted!") // HL
}

func partial() {
	// START OMIT
	fmt.Println("I am displayed partially!")
	// END OMIT
}
