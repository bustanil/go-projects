package main

import "fmt"
import "github.com/bustanil/oop/classes"

func main()  {
  var a [10]int
  a[0] = 1
  a[1] = 2

  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }

  strings := [3]string{"A", "B", "C"}

  for i := 0; i < len(strings); i++ {
    fmt.Println(strings[i])
  }
}
