package classes

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Age       int
	birthDate time.Time
}

func (p Person) Print() {
	fmt.Printf("Hi my name is %s and my age is %d\n", p.Name, p.Age)
}

func (p *Person) IncrementAge(age int) {
	p.Age = p.Age + age
}
