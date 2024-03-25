package classes

type AgeIncrementer interface {
	IncrementAge(increment int)
}

type Printer interface {
	Print()
}
