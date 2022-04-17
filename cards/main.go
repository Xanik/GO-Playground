package main

func main() {
	cards := newDeck()

	cards.shuffle()

	cards.print()

	cards.saveToFile("my_card")
}
