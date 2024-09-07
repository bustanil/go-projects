package main

import (
	"fmt"

	"github.com/bustanil/oop/classes"
)

func main() {
	bustanil := &classes.Person{
		Name: "Bustanil",
		Age:  40,
	}

	cat := classes.Animal{
		Name: "Black",
	}

	bustanil.Print()

	bustanil.IncrementAge(10)

	bustanil.Print()

	var ai classes.AgeIncrementer = bustanil
	var p classes.Printer = bustanil
	ai.IncrementAge(20)

	ai = cat
	fmt.Printf("%v, %T\n", ai, ai)

	bustanil.Print() // ok, copy happens in local scope of Print() function
	p.Print()

	var ai2 classes.AgeIncrementer = &classes.Person{
		Name: "Andi",
		Age:  30,
	}

	var bustanil2 *classes.Person = ai2.(*classes.Person)
	bustanil2.Print()

	incrementAge(ai, 5)
}

func incrementAge(ai classes.AgeIncrementer, value int) {
	ai.IncrementAge(value)
}

func print(p classes.Printer) {
	p.Print()
}
