package main

import "fmt"

// perimeter of rectangle is 2(l + w)
// area of rectangle is l * w


type Rectangle struct {
    x1 int32
    y1 int32
    x2 int32
    y2 int32
}

type Edge struct {
  l int32
  w int32
}


func calcEdge(coords Rectangle) Edge {
  edges := Edge{l: coords.x1 + coords.x2, w: coords.y1 + coords.y2}

  return edges
}

func area(edges Edge) {
  fmt.Println(edges.l * edges.w)
}

func perimeter(edges Edge) {
  fmt.Println(2 * (edges.l + edges.w))
}


func main() {
    rect1 := Rectangle{0, 7, 10, 0}
    edges := calcEdge(rect1)
    area(edges)
    perimeter(edges)
}
