package main

import "fmt"


type ServerStatus struct {
  Online string
}

type Server struct {
  Status string
}


func showServerStatus(servers map[string]Server) {
  for _, element := range servers {
    fmt.Println(element)
  }
}

func main() {

  servers := map[string]Server{
    "Devs Server": {Status: "Online"},
  }

  showServerStatus(servers)

}
