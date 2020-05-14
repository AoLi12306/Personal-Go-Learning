package main

import (
	"fmt"

	B "../basic"
	C "../concurrency"
)

func main() {

	B.Arr()
	B.Map()

	fmt.Println(B.Sum(1, 2))

	result, err := B.Sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	B.Stru()

	B.Pointer()

	C.Synccom()
	// C.Workpool()
}
