package main

import "fmt"

func main() {

	b := struct {
		color string
		size  int
	}{
		color: "red",
		size:  13,
	}

	fmt.Println(b)
}
