package basic

import "fmt"

type person struct {
	name string
	age  int
}

func Stru() {
	p := person{name: "David", age: 20}
	fmt.Println(p)
	fmt.Println(p.age)

}
