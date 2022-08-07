package main

import "fmt"

type Person struct {
   age int32
   name string
   weight float32
}

func main() {
  john := Person{age: 33, name: "fuckward", weight: 33.444}

  fmt.Println(john.age, john.name, john.weight)

}
