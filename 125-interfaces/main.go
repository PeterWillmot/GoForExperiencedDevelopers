package main

import (
	"fmt"
)

type Walker interface {
	Walkies(distance int)
}

type Person struct {
	PersonName string
}

func (p Person) Walkies(distance int) {
	fmt.Println(p.PersonName, "walking for", distance)
}

type Dog struct {
	PetName string
}

func (d Dog) Walkies(distance int) {
	fmt.Println(d.PetName, "walking for", distance)
}

func main() {
	human := Person{"Joe"}
	pup := Dog{"Dogbert"}

	GoWalkies(human, 3)
	GoWalkies(pup, 7)
}

func GoWalkies(subject Walker, distance int) {
	subject.Walkies(distance)
}
