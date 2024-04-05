package main

import "fmt"

type Muscle int

const (
	Unknown Muscle = iota
	Chest
	Shoulders
	Biceps
	Triceps
	Legs
)

func (m Muscle) String() string {
	switch m {
	case Chest:
		return "chest"
	case Shoulders:
		return "shoulders"
	case Biceps:
		return "biceps"
	case Triceps:
		return "triceps"
	case Legs:
		return "legs"
	}
	return "unknown"
}

func main() {
	fmt.Println(Unknown)
	fmt.Println(Chest)
	fmt.Println(Shoulders)
	fmt.Println(Biceps)
	fmt.Println(Triceps)
	fmt.Println(Legs)
}
