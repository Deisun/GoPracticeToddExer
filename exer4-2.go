package main

import "fmt"

func main() {

	s := []int{5, 6, 7, 8, 9}

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", s)

}
