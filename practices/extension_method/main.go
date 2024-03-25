package main

import "fmt"

type Number int

func (n Number) times(operand Number) Number {
  return n * Number(operand)
}

func (n Number) add(operand Number) Number {
  return n + Number(operand)
}

func main()  {
  fmt.Println(Number(1).times(2))
  fmt.Println(Number(1).add(2))

}
