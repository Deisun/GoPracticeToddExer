package main

import "fmt"

func main() {

	a := [5]int{22, 23, 24, 25, 26}

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", a)
}
