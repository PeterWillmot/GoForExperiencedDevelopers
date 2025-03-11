package main

import (
	"fmt"
)

func main() {
	typesOfThings := make(map[string]string)
	typesOfThings["dog"] = "animal"
	typesOfThings["cat"] = "animal"
	typesOfThings["tree"] = "plant"

	myType, ok := typesOfThings["dog"]
	fmt.Println(ok, myType) // true  animal

	myType, ok = typesOfThings["widget"]
	fmt.Println(ok, myType) // false

}
