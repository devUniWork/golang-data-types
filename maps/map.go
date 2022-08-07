package main

import "fmt"


func iterateMap(myMap map[string]int) {
  for _, value := range myMap {
    fmt.Println(value)
  }
}

func main() {
 myMap := map[string]int{
   "item 1": 1,
   "item 2": 2,
 }

//  fmt.Println(myMap["item 1"])
//
//  myMap["item 3"] = 6
//  fmt.Println(myMap)
//  delete(myMap, "item 1")
//  fmt.Println(myMap)

 iterateMap(myMap)

}
