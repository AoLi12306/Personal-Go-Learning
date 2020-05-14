package basic

import (
	"fmt"
)

func Arr() {
	a := []int{5, 4, 3, 2, 1}
	a = append(a, 13)

	for index, value := range a {
		fmt.Println(index, value)
	}
}

func Map() {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "dodecagon")

	for key, value := range vertices {
		fmt.Println(key, value)
	}
}
