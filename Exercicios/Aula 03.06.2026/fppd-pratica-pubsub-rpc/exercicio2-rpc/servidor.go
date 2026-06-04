package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"net/rpc"
)

// --- Tipos dos argumentos e respostas RPC ---

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

// --- Serviço RPC ---

type NumericaService struct{}

func ehPrimo(n int64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func (s *NumericaService) ContarPrimos(args *Intervalo, reply *int64) error {
	var count int64
	for n := args.Inicio; n <= args.Fim; n++ {
		if ehPrimo(n) {
			count++
		}
	}
	*reply = count
	fmt.Printf("  ContarPrimos(%d, %d) = %d\n", args.Inicio, args.Fim, *reply)
	return nil
}

func (s *NumericaService) EstimarPi(args *ArgIteracoes, reply *ResultadoPi) error {
	var dentro int64
	r := rand.New(rand.NewSource(42))
	for i := int64(0); i < args.N; i++ {
		x, y := r.Float64(), r.Float64()
		if x*x+y*y <= 1.0 {
			dentro++
		}
	}
	reply.Valor = math.Round(4.0*float64(dentro)/float64(args.N)*1e6) / 1e6
	reply.Amostras = args.N
	fmt.Printf("  EstimarPi(%d) = %f\n", args.N, reply.Valor)
	return nil
}

// TODO: implemente SomarIntervalo no servidor.
// Assinatura: func (s *NumericaService) SomarIntervalo(args *Intervalo, reply *int64) error
// A função deve somar todos os inteiros de args.Inicio até args.Fim (inclusive).

func main() {
	endereco := ":9000"

	rpc.Register(new(NumericaService))

	listener, err := net.Listen("tcp", endereco)
	if err != nil {
		log.Fatalf("Erro ao abrir porta %s: %v\n", endereco, err)
	}
	defer listener.Close()

	fmt.Printf("Servidor RPC escutando em %s\n", endereco)
	fmt.Println("Aguardando chamadas... (Ctrl+C para encerrar)")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão: %v\n", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
