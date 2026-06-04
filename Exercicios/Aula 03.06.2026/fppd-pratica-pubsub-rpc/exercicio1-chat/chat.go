package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsURL       = "nats://demo.nats.io:4222"
	subjectPrefix = "fppd"
)

func chatSubject(sala string) string {
	return fmt.Sprintf("%s.%s.chat", subjectPrefix, sala)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Uso: go run chat.go <sala> <nome>\n")
		fmt.Fprintf(os.Stderr, "  sala: identificador da sala (ex: turma32)\n")
		fmt.Fprintf(os.Stderr, "  nome: seu nome no chat\n")
		os.Exit(1)
	}

	sala := os.Args[1]
	nome := os.Args[2]
	subject := chatSubject(sala)

	// Conecta ao servidor NATS
	nc, err := nats.Connect(natsURL,
		nats.Name(fmt.Sprintf("fppd-chat-%s", nome)),
		nats.ReconnectWait(2*time.Second),
		nats.MaxReconnects(10),
	)
	if err != nil {
		log.Fatalf("Erro ao conectar ao NATS (%s): %v\n", natsURL, err)
	}
	defer nc.Close()

	fmt.Printf("Conectado ao servidor NATS: %s\n", natsURL)
	fmt.Printf("Sala: %s | Subject: %s\n", sala, subject)
	fmt.Printf("Seu nome: %s\n", nome)
	fmt.Println(strings.Repeat("─", 50))

	// Inscreve-se no subject da sala para receber mensagens do grupo
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		texto := string(msg.Data)
		partes := strings.SplitN(texto, "|", 3)
		if len(partes) == 3 {
			remetente := partes[0]
			hora := partes[1]
			conteudo := partes[2]
			if remetente != nome {
				fmt.Printf("[%s] %s: %s\n", hora, remetente, conteudo)
			}
		}
	})
	if err != nil {
		log.Fatalf("Erro ao se inscrever no subject %s: %v\n", subject, err)
	}

	// Publica mensagem de entrada na sala
	entradaMsg := fmt.Sprintf("%s|%s|entrou na sala", nome, time.Now().Format("15:04:05"))
	nc.Publish(subject, []byte(entradaMsg))

	// Tratamento de Ctrl+C — publica saída antes de desconectar
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		saidaMsg := fmt.Sprintf("%s|%s|saiu da sala", nome, time.Now().Format("15:04:05"))
		nc.Publish(subject, []byte(saidaMsg))
		nc.Flush()
		fmt.Println("\nDesconectado.")
		os.Exit(0)
	}()

	// Loop de leitura do teclado
	fmt.Println("Digite suas mensagens (Enter para enviar, Ctrl+C para sair):")
	fmt.Println(strings.Repeat("─", 50))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		if texto == "" {
			continue
		}

		mensagem := fmt.Sprintf("%s|%s|%s", nome, time.Now().Format("15:04:05"), texto)
		err := nc.Publish(subject, []byte(mensagem))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao enviar: %v\n", err)
			continue
		}
		fmt.Printf("[%s] você: %s\n", time.Now().Format("15:04:05"), texto)
	}
}
