package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Testando conexão com demo.nats.io...")

	nc, err := nats.Connect("nats://demo.nats.io:4222",
		nats.Timeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Falha na conexão: %v\n", err)
	}
	defer nc.Close()

	fmt.Printf("Conectado! Servidor: %s\n", nc.ConnectedUrl())

	subject := "fppd.teste.ping"

	recebido := make(chan bool, 1)
	nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("Recebido: %s\n", string(msg.Data))
		recebido <- true
	})

	nc.Publish(subject, []byte("hello from FPPD!"))
	nc.Flush()

	select {
	case <-recebido:
		fmt.Println("Pub/Sub funcionando corretamente!")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout esperando mensagem (pode haver firewall)")
	}
}
