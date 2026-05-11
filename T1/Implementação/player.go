// Funcionamento desejado:
// Cria uma struct para definir o jogador (nome, id, bool para verificar se é bot)
// Cria uma struct para definir os comandos, ou seja, uma ação enviada para o jogo (“Don't communicate by sharing memory; share memory by communicating.”)
// Função simples para criar os players, sendo um jogador e dois bots
// Função que cuida dos comportamentos automáticos para os bots com o seguinte funcionamento:
// 	Recebe player Player, ou seja, o bot que está exercutando
// 	Recebe chan <- Command, canal somente para envios
// 	Recebe done <- chan struct{}, canal somente para recepção
//  Roda um loop "infinito", e abre um select, vê se o jogo continua ou acabou
// 	Bota um tempo de reação pro bot jogar, fica em torno de 0,5s - 2s
//	Envia a ação do bot para o canal

package main

import (
	"math/rand"
	"time"
)

type Player struct {
	name string
	id   int
	bot  bool
	hand []Card
}

type Command struct {
	playerId int
	action   string
}

func createPlayers() []Player {
	return []Player{
		{id: 1, name: "Jogador", bot: false},
		{id: 2, name: "Bot 1", bot: true},
		{id: 3, name: "Bot 2", bot: true},
	}
}

func botPlayer(player Player, commandCh chan<- Command, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return

		default:
			time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)

			commandCh <- Command{
				playerId: player.id,
				action:   "slap",
			}
		}
	}
}
