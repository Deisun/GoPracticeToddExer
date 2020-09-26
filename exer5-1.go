package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	flavors   []string
}

func main() {
	p1 := person{
		firstname: "Rob",
		lastname:  "Douma",
		flavors:   []string{"chocolate", "vanilla"},
	}

	p2 := person{
		firstname: "Tania",
		lastname:  "Ocasio",
		flavors:   []string{"rocky roads", "mint chocolate chip"},
	}

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for _, v := range p1.flavors {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
