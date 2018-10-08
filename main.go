package main

import "fmt"

func main() {
	var card string = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
