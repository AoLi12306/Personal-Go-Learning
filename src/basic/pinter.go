package basic

import "fmt"

func Pointer() {
	i := 7

	inc2(i)
	fmt.Println(i)

	inc(&i)
	fmt.Println(i)

}

func inc(x *int) {
	*x++
}

func inc2(x int) {
	x++
}
