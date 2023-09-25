package main

import "fmt"

func main() {

	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = taxRate * float64(price)
	fmt.Println("Tax is", tax, "dollars.")

	var total int = price + int(tax)
	fmt.Println("Total cost is", total, "dollars")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available")
	fmt.Println("Within budget?")

}
