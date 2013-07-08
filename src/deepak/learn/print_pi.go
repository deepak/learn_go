# http://tour.golang.org/#3
# QUESTION: Programs start running in package main.
# QUESTION: By convention, the package name is the same as the last element of the import path.
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Happy", math.Pi, "Day")
}