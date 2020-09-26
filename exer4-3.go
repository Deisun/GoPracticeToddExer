package main

import "fmt"

func main() {

	s := []int{3, 4, 5, 6, 7}

	s = append(s, 8)

	fmt.Println(s)

	t := []int{9, 10, 11}

	s = append(s, t...)

	fmt.Println(s)

}
