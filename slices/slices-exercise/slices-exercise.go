package main

import "fmt"


type Part struct {
  Name string
  Weight int
}


func printAssemblyLine(parts []Part) {
  for i := 0; i < len(parts); i++ {
    part := parts[i]
    fmt.Println(part)
  }
}

func main() {
 parts := []Part{
   {Name: "lid", Weight: 10},
   {Name: "body", Weight: 1},
   {Name: "trunk", Weight: 4},
 }
 printAssemblyLine(parts)

 newParts := []Part{
  {Name: "window", Weight: 3},
  {Name: "Side panel", Weight: 3},
 }

 parts = append(parts, newParts...)

 fmt.Println()
 printAssemblyLine(parts)
 fmt.Println()
 printAssemblyLine(parts[3:])


}
