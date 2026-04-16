// Exercício 2B — Largada Sincronizada ("Ready-Set-Go")
//
// Este programa simula uma corrida: 6 corredores devem esperar
// até que o juiz dê o sinal de largada. O juiz prepara a pista
// (simula com sleep) e então libera todos de uma vez.
//
// Atualmente NÃO há sincronização: os corredores partem
// imediatamente, sem esperar o juiz.
//
// Sua tarefa: usar sync.Cond com Broadcast() para que todos
// os corredores esperem o sinal do juiz antes de começar.

//Pergunta: os corredores esperam o juiz? Qual é o problema com a saída atual?
//Não esperam, estão saindo antes.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numCorredores = 6

var (
	mu     sync.Mutex
	cond   = sync.NewCond(&mu)
	pronto bool
)

func corredor(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Corredor %d] na linha de largada\n", id)

	// TODO: esperar o sinal de largada
	mu.Lock()
	for !pronto {
		cond.Wait()
	}
	mu.Unlock()

	duracao := time.Duration(500+rand.Intn(2000)) * time.Millisecond
	fmt.Printf("[Corredor %d] LARGOU!\n", id)
	time.Sleep(duracao)
	fmt.Printf("[Corredor %d] chegou (tempo: %v)\n", id, duracao)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= numCorredores; i++ {
		wg.Add(1)
		go corredor(i, &wg)
	}

	// Juiz prepara a pista
	fmt.Println("\n[Juiz] Preparando a pista...")
	time.Sleep(2 * time.Second)

	// TODO: sinalizar a largada
	mu.Lock()
	pronto = true
	cond.Broadcast()
	mu.Unlock()

	fmt.Println("[Juiz] VAI!\n")

	wg.Wait()
	fmt.Println("\nCorrida encerrada.")
}

/*
Perguntas para reflexão
1. Por que usamos Broadcast() e não Signal()? O que aconteceria se usássemos
Signal()?
	Se usássemos Signal(), apenas um corredor seria acordado.

2. Por que usamos for !pronto (loop) e não if !pronto? O que poderia dar errado
com if?
	Se usássemos if, poderia continuar mesmo sem estar realmente liberada, devido ao despertar
	indevido ou pois outra goroutine alterou o estado antes ela reassumir o lock.

3. O que aconteceria se o juiz chamasse Broadcast() sem alterar pronto para
true?
	Todos seriam acordadois, mas devido a pronto continuar como false, voltariam a esperar.
*/
