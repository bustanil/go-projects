package main

import "fmt"

func main() {
	// literals are untyped (default integer literal is int)
	var anInteger = 1_000_000_000_000_000
	fmt.Println(anInteger)

	var anotherIntegerVar = anInteger
	fmt.Println(anotherIntegerVar) 
	// because anInteger is an int. then anotherIntegerVar is also an int

	// assign literal to a typed variable
	var anotherInteger int = 1_000_000_000_000_000
	fmt.Println(anotherInteger)

	// convert untyped literal to typed
	var aVeryBigInteger = float64(1_000_000_000_000_000_000_000_000)
	fmt.Println(aVeryBigInteger)

	// rune
	var aRune = 'a'
	var aUnicode = '\u1F60'
	fmt.Println(aRune, aUnicode)

	// string literal
	var aString = "bustanil"
	fmt.Println(aString)

	// boolean literal
	var aBoolean = true
	fmt.Println(aBoolean)

	// float literal (default float64)
	var aFloat64 = 1.1
	fmt.Println(aFloat64)
}
