package main

import "fmt"

func main() {
	cow := Cow{
		Liters:  12,
		Name:    "Berta",
		Spots:   4,
		Meat:    "Beef",
		Species: "Cow",
	}

	fmt.Println(cow)
}

type Cow struct {
	Liters  int
	Name    string
	Spots   int
	Meat    string
	Species string
}
