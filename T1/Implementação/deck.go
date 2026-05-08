// Funcionamento desejado:
// Cria as cartas como constantes do tipo "Card".
// Cria uma listinha com todas as cartas.
// Em 2 for encadeados, cria o deck de 36 cartas, sendo 12 para cadaa jogador, de maneria aleatória.
// Embaralhar de cartas utilizando uma função anônima para trocar as cartas de posições.
// Este arquivo não implementa nenhum conceito de concorrência, somente implementa a lógica das cartas.

package main

import "math/rand"

type Card string

const Taco Card = "Taco"
const Gato Card = "Gato"
const Cabra Card = "Cabra"
const Queijo Card = "Queijo"
const Pizza Card = "Pizza"

func NewDeck() []Card {
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
