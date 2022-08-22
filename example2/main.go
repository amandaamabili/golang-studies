package main

import (
	"fmt"
	"github.com/amandaamabili/golang-studies/src/product"
)

func main() {
	p := product.Product{
		Id:    "1d232",
		Name:  "Rossini",
		Price: 5.5,
	}
	fmt.Println(p)
}
