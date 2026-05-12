package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func renderer(renderCh <-chan GameEvent, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("\nEncerrando renderização...")
			return

		case event := <-renderCh:
			fmt.Print("\033[H\033[2J")

			fmt.Println("===================================")
			fmt.Println("      Taco Gato Cabra Queijo Pizza")
			fmt.Println("===================================")
			fmt.Println()

			fmt.Println(event.message)

			fmt.Println()
			fmt.Println("Comandos:")
			fmt.Println("ENTER -> revelar carta")
			fmt.Println("s     -> bater no monte")
			fmt.Println("q     -> sair")
		}
	}
}

func inputPlayer(commandCh chan<- Command, done chan struct{}) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')

		select {
		case <-done:
			return

		default:
			switch input[0] {

			case 's':
				commandCh <- Command{
					playerId: 1,
					action:   "slap",
				}

			case 'q':
				close(done)
				return

			default:
				commandCh <- Command{
					playerId: 1,
					action:   "reveal",
				}
			}
		}
	}
}

func automaticTurns(commandCh chan<- Command, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return

		default:
			time.Sleep(2 * time.Second)

			commandCh <- Command{
				playerId: 0,
				action:   "reveal",
			}
		}
	}
}

func main() {
	commandCh := make(chan Command)
	renderCh := make(chan GameEvent)
	done := make(chan struct{})

	players := createPlayers()

	go gameLoop(commandCh, renderCh, done)

	go renderer(renderCh, done)

	go inputPlayer(commandCh, done)

	go automaticTurns(commandCh, done)

	for _, player := range players {
		if player.bot {
			go botPlayer(player, commandCh, done)
		}
	}

	<-done

	time.Sleep(4000 * time.Millisecond)

	fmt.Println("\nFim de jogo.")
}
