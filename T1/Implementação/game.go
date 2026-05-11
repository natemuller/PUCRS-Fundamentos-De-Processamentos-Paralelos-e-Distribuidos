// Funcionamento desejado:
// Define os eventos enviados do jogo para a renderização.
// Define o estado central da partida, incluindo deck, jogadores, monte central,
// turno atual, sequência do jogo, controle de match e fim de jogo.
// A função gameLoop concentra a lógica principal da partida.
// Ela recebe comandos pelo canal commandCh, envia mensagens pelo canal renderCh
// e escuta o canal done para encerrar de forma ordenada.
// O select dentro do gameLoop multiplexa eventos de comando e encerramento.
// Apenas o gameLoop altera o estado do jogo, evitando que bots ou input
// modifiquem diretamente deck, mãos dos jogadores ou regras da partida.
// A função playTurn executa uma rodada, revelando carta da mão do jogador atual.
// Quando há match, libera a possibilidade de bater no monte.
// A função handleSlap trata as batidas no monte e aplica a penalidade.

package main

import "fmt"

type GameEvent struct {
	message string
}

type GameState struct {
	players       []Player
	pile          []Card
	currentTurn   int
	sequenceIndex int
	canSlap       bool
	slapOrder     []int
	finished      bool
}

var sequence = []Card{Taco, Gato, Cabra, Queijo, Pizza}

func gameLoop(commandCh <-chan Command, renderCh chan<- GameEvent, done <-chan struct{}) {
	deck := newDeck()
	players := createPlayers()
	players = dealCards(deck, players)

	state := GameState{
		players: players,
	}

	renderCh <- GameEvent{message: "Jogo iniciado!"}

	for {
		select {
		case <-done:
			renderCh <- GameEvent{message: "Encerrando game loop..."}
			return

		case cmd := <-commandCh:
			if state.finished {
				continue
			}

			if cmd.action == "reveal" {
				playTurn(&state, renderCh)
			}

			if cmd.action == "slap" {
				handleSlap(&state, cmd, renderCh)
			}
		}
	}
}

func playTurn(state *GameState, renderCh chan<- GameEvent) {
	if state.canSlap {
		renderCh <- GameEvent{
			message: "Match pendente. Todos precisam bater antes da próxima carta.",
		}
		return
	}

	player := &state.players[state.currentTurn]

	if len(player.hand) == 0 {
		renderCh <- GameEvent{
			message: fmt.Sprintf("%s não tem mais cartas. Fim de jogo!", player.name),
		}
		state.finished = true
		return
	}

	card := player.hand[0]
	player.hand = player.hand[1:]

	if len(player.hand) == 0 {
		renderCh <- GameEvent{
			message: fmt.Sprintf("%s venceu o jogo!", player.name),
		}
		state.finished = true
		return
	}

	state.pile = append(state.pile, card)

	word := sequence[state.sequenceIndex]

	renderCh <- GameEvent{
		message: fmt.Sprintf("%s revelou: %s | Palavra dita: %s", player.name, card, word),
	}

	if card == word {
		state.canSlap = true
		state.slapOrder = []int{}

		renderCh <- GameEvent{
			message: "MATCH! Todos devem bater no monte!",
		}
	} else {
		state.canSlap = false
	}

	state.sequenceIndex = (state.sequenceIndex + 1) % len(sequence)
	state.currentTurn = (state.currentTurn + 1) % len(state.players)
}

func handleSlap(state *GameState, cmd Command, renderCh chan<- GameEvent) {
	if !state.canSlap {
		renderCh <- GameEvent{
			message: fmt.Sprintf("Jogador %d bateu fora de hora.", cmd.playerId),
		}
		return
	}

	for _, id := range state.slapOrder {
		if id == cmd.playerId {
			return
		}
	}

	state.slapOrder = append(state.slapOrder, cmd.playerId)

	renderCh <- GameEvent{
		message: fmt.Sprintf("Jogador %d bateu no monte!", cmd.playerId),
	}

	if len(state.slapOrder) == len(state.players) {
		loserId := state.slapOrder[len(state.slapOrder)-1]

		for i := 0; i < len(state.players); i++ {
			if state.players[i].id == loserId {
				state.players[i].hand = append(state.players[i].hand, state.pile...)

				renderCh <- GameEvent{
					message: fmt.Sprintf("%s foi o último e pegou o monte!", state.players[i].name),
				}

				break
			}
		}

		state.pile = []Card{}
		state.canSlap = false
		state.slapOrder = []int{}
	}
}
