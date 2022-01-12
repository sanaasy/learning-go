package main

import "fmt"

var deckSize int

func main() {
	// var card string = "Ace of Spades"
	oldCard := "Ace of Spades" // to initalize
	oldCard = "Five of Diamonds"

	deckSize = 52

	fiveOfDiamonds := newCard()

	// a slice has to have the same data type
	cards := []string{"Ace of Diamonds", newCard()}
	// creates a new slice and sets it to cards
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(deckSize)
	fmt.Println(fiveOfDiamonds)
	fmt.Println(oldCard)
}

func newCard() string {
	return "Five of Diamonds"
}
