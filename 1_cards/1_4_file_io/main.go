package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	newCard := newDeckFromFile("my_cards")

	newCard.print()

	errorCard := newDeckFromFile("does_not_exist")
	errorCard.print()
}
