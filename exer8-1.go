package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Firstname string
	Lastname  string
}

func main() {
	u1 := user{
		Firstname: "Rob",
		Lastname:  "Douma",
	}

	u2 := user{
		Firstname: "Tania",
		Lastname:  "Ocasio",
	}

	users := []user{u1, u2}

	j, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))

}
