package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// --- Tipos compartilhados (devem ser iguais aos do servidor) ---

type Intervalo struct {
	Inicio, Fim int64
}

type ArgIteracoes struct {
	N int64
}

type ResultadoPi struct {
	Valor    float64
	Amostras int64
}

func main() {
	endereco := "localhost:9000"
	if len(os.Args) >= 2 {
		endereco = os.Args[3]
	}

	fmt.Printf("Conectando ao servidor RPC em %s...\n", endereco)
	client, err := rpc.Dial("tcp", endereco)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v\n", err)
	}
	defer client.Close()
	fmt.Println("Conectado!\n")

	// --- ContarPrimos ---
	var primos int64
	err = client.Call("NumericaService.ContarPrimos", &Intervalo{1, 100_000}, &primos)
	if err != nil {
		log.Printf("Erro em ContarPrimos: %v\n", err)
	} else {
		fmt.Printf("ContarPrimos(1, 100000) = %d\n", primos)
	}

	// --- EstimarPi ---
	var pi ResultadoPi
	err = client.Call("NumericaService.EstimarPi", &ArgIteracoes{1_000_000}, &pi)
	if err != nil {
		log.Printf("Erro em EstimarPi: %v\n", err)
	} else {
		fmt.Printf("EstimarPi(%d)           = %f\n", pi.Amostras, pi.Valor)
	}

	// --- SomarIntervalo ---
	var soma int64
	err = client.Call("NumericaService.SomarIntervalo", &Intervalo{1, 1000}, &soma)
	if err != nil {
		log.Printf("Erro em SomarIntervalo: %v\n", err)
	} else {
		fmt.Printf("SomarIntervalo(1, 1000) = %d\n", soma)
	}
}
