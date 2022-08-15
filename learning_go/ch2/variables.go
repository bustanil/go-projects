package main

import "fmt"

func main() {
	// there are many types of int (based on size)
	var aUint16 uint16 = 100
	var aUint32 uint32 = 200

	fmt.Println(aUint16, aUint32)

	// but for most cases int or uint is sufficient
	var aInt int = 1
	var aUint uint = 1_000

	fmt.Println(aInt, aUint)

	// "zero values"
	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroBoolean bool  // false
	var zeroString string // empty string

	fmt.Println(zeroInt, zeroFloat, zeroBoolean, zeroString)

	// declare and init multiple variables
	var someInt, someFloat, someBool = 1, 80.5, true
	fmt.Println(someInt, someFloat, someBool)

	// declare and assign var with :=
	name := "bustanil"
	fmt.Println(name)
	// declare and assign multiple variables
	name, address := "john", "new york"
	fmt.Println(name, address)
	// NOTE: this is usually only used within functions

	// multiple variable declarations
	var (
		anotherInt                      int  = 19
		anotherBool                     bool = false
		anotherString, andAnotherString      = "test", "test2"
		anotherFloat                         = 99.9
	)

	fmt.Println(anotherInt, anotherBool, anotherString, andAnotherString, anotherFloat)
}
