package main
import "fmt"

type Product struct {
  Name string
  Price float32
}

func printLastItem(product Product) {
 fmt.Println(product.Name, "is the last shopping list item")
}

func totalCost(cost float32) {
 fmt.Println("total cost is", cost)
}

func totalNumberItems(items int) {
 fmt.Println("total number of items is", items)
}

func checkPrice(products [3]Product) {
  var cost float32
  var numberItems int
  for i := 0; i < len(products); i++ {
    product := products[i]
    j := i

    if (j == len(products) - 1) {
        printLastItem(product)
    }
    cost = cost + product.Price
    numberItems++
  }
  totalCost(cost)
  totalNumberItems(numberItems)
}


func main() {

 shoppingList := [...]Product{
    {Name: "tampons", Price: 5.00},
    {Name: "cheese", Price: 6.00},
    {Name: "Basil", Price: 3.00},
 }

 checkPrice(shoppingList)

}
