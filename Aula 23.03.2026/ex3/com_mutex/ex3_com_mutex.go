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
		Contador: 200000 (esperado: 200000) - SEM -race
		Contador: 200000 (esperado: 200000) - COM -race
		Sempre 200000!

4. Compare o tempo de execução com e sem mutex
		Com Mutex: 3.2398ms.
		Com mutex, o acesso à variável compartilhada é sincronizado, garantindo o valor correto de 200000 em todas as execuções,
		embora com um pequeno aumento no tempo de execução por causa do custo de bloqueio e desbloqueio da região crítica.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var contador int
var mu sync.Mutex

func incrementar(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		mu.Lock()
		contador++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	inicio := time.Now()

	go incrementar(100000, &wg)
	go incrementar(100000, &wg)

	wg.Wait()
	duracao := time.Since(inicio)

	fmt.Printf("Contador: %d (esperado: 200000)\n", contador)
	fmt.Printf("Tempo com mutex: %v\n", duracao)
}
