package main

import "fmt"

type student struct {
	firstname string
	lastname  string
}

func (s *student) changeMe() {
	s.firstname = "Blah"
}

func changeBe(s *student) {
	s.lastname = "Flah"
}

func main() {

	s1 := student{
		firstname: "Rob",
		lastname:  "Douma",
	}

	fmt.Println(s1)
	s1.changeMe()
	fmt.Println(s1)
	changeBe(&s1)
	fmt.Println(s1)

}
