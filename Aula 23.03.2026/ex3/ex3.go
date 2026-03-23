/*
Objetivo: identificar uma condição de corrida e usar o detector de corrida do Go.
1. Execute o programa 5 vezes com  go run ex3.go . O resultado é sempre 200000? Anote os valores
   obtidos.
   		100811, 100975, 100847, 100000, 102710.

2. Execute com o detector de corrida:  go run -race ex3.go . O que a saída indica?
		Contador: 196722 (esperado: 200000)
		Found 2 data race(s)
		exit status 66

3. Corrija o programa adicionando um  sync.Mutex  para proteger o acesso ao contador. Execute
   novamente (com e sem  -race ) e verifique que o resultado é sempre 200000.
		Não deu tempo de fazer...

4. Compare o tempo de execução com e sem mutex
*/

package main

import (
	"fmt"
	"sync"
)

var contador int

func incrementar(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		contador++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementar(100000, &wg)
	go incrementar(100000, &wg)

	wg.Wait()
	fmt.Printf("Contador: %d (esperado: 200000)\n", contador)
}
