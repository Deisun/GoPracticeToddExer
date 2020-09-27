package main

import "fmt"

type teacher struct {
	firstname string
	lastname  string
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (t *teacher) speak() {
	fmt.Printf("%v would like to have a word\n", t.firstname)
}

func main() {

	h1 := teacher{
		firstname: "Rob",
	}

	h2 := teacher{
		firstname: "Tania",
	}

	saySomething(&h1)
	saySomething(&h2)

}
