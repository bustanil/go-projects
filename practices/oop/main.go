package main

import (
	"github.com/bustanil/oop/classes"
)

func main() {
	bustanil := &classes.Person{
		Name: "Bustanil",
		Age:  40,
	}

	bustanil.Print()

	bustanil.IncrementAge(10)

	bustanil.Print()

	var ai classes.AgeIncrementer = bustanil
	var p classes.Printer = bustanil
	ai.IncrementAge(20)

	bustanil.Print() // ok, copy happens in local scope of Print() function
	p.Print()

	var ai2 classes.AgeIncrementer = &classes.Person{
		Name: "Andi",
		Age:  30,
	}

	var bustanil2 *classes.Person = ai2.(*classes.Person)
	bustanil2.Print()
}
