package main

import (
  "fmt"
  "time"
  "math/rand"
)

func roll(sides int) int {
  return rand.Intn(sides) + 1
}


func main() {
  rand.Seed(time.Now().UnixNano())

  dice, sides := 2, 12
  rolls := 9999999

  for r := 1; r <= rolls; r++ {
    sum := 0

    for d := 1; d <= dice; d++ {
        rolled := roll(sides)
        sum += rolled
        fmt.Println("Roll #", r, ", Die #", d, ":", rolled)
    }

    fmt.Println("Total rolled:", sum)
    switch sum := sum; {
    case sum == 2 && dice == 2:
    fmt.Println("Snake Eyes!")
    case sum == 7:
    fmt.Println("Lucky 7!")
    case sum%2 == 0:
    fmt.Println("Even!")
    case sum%2 == 1:
    fmt.Println("Odd!")

    }

  }

//   fmt.Println(rand.Intn(6))

}
