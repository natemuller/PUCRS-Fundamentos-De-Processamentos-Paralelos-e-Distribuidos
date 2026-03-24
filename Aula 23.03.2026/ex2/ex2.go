/*
Objetivo: criar goroutines, sincronizá-las com  sync.WaitGroup  e proteger dados compartilhados com
sync.Mutex .
Crie um programa em Go ( ex2.go ) com duas goroutines disparadas a partir da  main() :
1. Uma goroutine deve imprimir números de 1 a 10, um a cada 1 segundo
2. A outra goroutine deve imprimir números de 10 a 1, um a cada 1 segundo
3. A main() deve esperar até que as duas goroutines terminem
4. Utilize sync.WaitGroup para fazer a main() esperar
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func imprimeDeUmADez(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		mu.Lock()
		fmt.Printf("%d ", i)
		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func imprimeDeDezAUm(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i >= 1; i-- {
		mu.Lock()
		fmt.Printf("%d ", i)
		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go imprimeDeUmADez(&wg)
	go imprimeDeDezAUm(&wg)
	wg.Wait()
}
