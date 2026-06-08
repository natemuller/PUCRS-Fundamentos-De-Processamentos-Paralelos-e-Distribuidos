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

	// Inscreve-se no subject privado para receber mensagens diretas
	dmSubject := fmt.Sprintf("%s.%s.dm.%s", subjectPrefix, sala, nome)
	_, err = nc.Subscribe(dmSubject, func(msg *nats.Msg) {
		texto := string(msg.Data)
		partes := strings.SplitN(texto, "|", 3)
		if len(partes) == 3 {
			remetente := partes
			hora := partes[1]
			conteudo := partes[2]
			// Formatação exigida para mensagens privadas recebidas
			fmt.Printf("[%s] (privado) %s: %s\n", hora, remetente, conteudo)
		}
	})
	if err != nil {
		log.Fatalf("Erro ao se inscrever no subject privado %s: %v\n", dmSubject, err)
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

		horaAtual := time.Now().Format("15:04:05")

		// Verifica se a mensagem é um comando privado usando as dicas do enunciado
		if strings.HasPrefix(texto, "/privado ") {
			// Divide em "/privado", "destinatario", "a mensagem..."
			partes := strings.SplitN(texto, " ", 3)

			if len(partes) == 3 {
				destinatario := partes[1]
				conteudo := partes[2]

				// Monta o subject do destinatário
				destSubject := fmt.Sprintf("%s.%s.dm.%s", subjectPrefix, sala, destinatario)
				mensagem := fmt.Sprintf("%s|%s|%s", nome, horaAtual, conteudo)

				err := nc.Publish(destSubject, []byte(mensagem))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Erro ao enviar mensagem privada: %v\n", err)
					continue
				}
				fmt.Printf("[%s] você (privado para %s): %s\n", horaAtual, destinatario, conteudo)
			} else {
				fmt.Println("Formato inválido. Use: /privado <destinatario> <mensagem>")
			}
		} else {
			// Comportamento normal para mensagens públicas na sala
			mensagem := fmt.Sprintf("%s|%s|%s", nome, horaAtual, texto)
			err := nc.Publish(subject, []byte(mensagem))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao enviar: %v\n", err)
				continue
			}
			fmt.Printf("[%s] você: %s\n", horaAtual, texto)
		}
	}
}
