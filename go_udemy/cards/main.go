package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spade"
	// card := "Ace of Spade"
	// cards := []string{"Ace of Hearts", newCard()}
	// cards := deck{"Ace of Hearts", newCard()}
	// cards = append(cards, "King of Spade")
	cards := newDeck()
	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	fmt.Println("Deck:")
	// cards.print()

	hand, remainingCards := cards.deal(5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println(len(remainingCards))
	fmt.Println("CARDS:")
	fmt.Println(len(cards))

	fmt.Println(cards.toString())

	hand.saveFile("hand1")

	//deck from file
	fmt.Println("Hand from File:")
	fromFile := newDeckFromFile("hand1")
	fromFile.print()

	//shuffle
	fromFile.shuffle()
	fmt.Println("SHUFFLED FILE")
	fromFile.print()
}
