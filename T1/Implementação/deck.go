// Funcionamento desejado:
// Cria as cartas como constantes do tipo "Card".
// Cria um slice com todos os tipos de cartas.
// Em dois for encadeados, cria o deck com 36 cartas,
// sendo 12 cartas para cada jogador.
// Embaralha as cartas utilizando uma função anônima
// para trocar cartas de posição.
// Este arquivo não implementa concorrência,
// apenas a lógica das cartas.

package main

import "math/rand"

type Card string

const Taco Card = "Taco"
const Gato Card = "Gato"
const Cabra Card = "Cabra"
const Queijo Card = "Queijo"
const Pizza Card = "Pizza"

func newDeck() []Card {
	deck := []Card{}

	cards := []Card{Taco, Gato, Cabra, Queijo, Pizza}

	for len(deck) < 36 {
		for _, card := range cards {
			if len(deck) == 36 {
				break
			}
			deck = append(deck, card)
		}
	}
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func dealCards(deck []Card, players []Player) []Player {
	cardIndex := 0

	for i := 0; i < len(players); i++ {
		for j := 0; j < 12; j++ {
			players[i].hand = append(players[i].hand, deck[cardIndex])
			cardIndex++
		}
	}

	return players
}
