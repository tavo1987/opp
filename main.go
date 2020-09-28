package main

import (
	"fmt"
)

func main() {
	golang := Course{
		Name:    "Golang from scratch",
		Price:   29.99,
		IsFree:  false,
		UserIDs: []uint{1, 6, 3, 5},
		Lessons: map[uint]string{
			1: "Introduction",
			2: "Structures",
			3: "Maps",
		},
	}

	golang.updatePrice(45.99)

	fmt.Println(golang)
}