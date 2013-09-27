package main

import "fmt"

func split(sum int) (x, y int, z string) {
  x = sum * 4 / 9
  y = sum - x
  z = "testing"
  return //multiple result parameters 
}

func main() {
  // QUESTION: do not care that x, y are not used. z is being used and that is what i want
  x, y, z := split(42)
  fmt.Println(z)
  // ./return_all.go:14: x declared and not used
  // ./return_all.go:14: y declared and not used
  // fmt.Println(x,y,z)
  fmt.Println(split(42))
}