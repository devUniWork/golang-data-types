package main

import "fmt"


func increment(value *int) int {
  return *value * 30
}

func main() {
 v := 10

 value := increment(&v)

 fmt.Println(value, v)

}
