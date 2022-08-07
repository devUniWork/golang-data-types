package main

import (
   "fmt"
   "strings"
)

type Person struct {
  Name string
  Age int
}

// this is not variadic... update example
func printNames(person []Person) string {
  for _, element := range person {
    si := element
    return strings.Join(si.Name, " ")
  }
  return si.Name
}

func main() {
    john := []Person{
    {Name: "simon", Age: 44},
    }

    fmt.Println(john)
}
