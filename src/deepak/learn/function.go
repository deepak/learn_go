package main

import "fmt"

func add(x int, y int) int {
  // can have semi-colons as line terminator. not required though
  // no implicit returns
  return x + y;
}

func main() {
  fmt.Println(add(1, 42))
}