FPPD — Prática: Publish-Subscribe e RPC em Go
===============================================

Este pacote contém os códigos-fonte de referência e pontos de partida
para o roteiro de prática sobre comunicação em grupo (pub-sub) e RPC.

Estrutura:
  exercicio1-chat/           Exercício 1 — Chat distribuído com NATS
    chat.go                  Chat funcional — executar e estender
    teste-conexao.go         Script de teste de conectividade

  exercicio2-rpc/            Exercício 2 — RPC com net/rpc
    servidor.go              Servidor RPC (adicionar SomarIntervalo)
    cliente.go               Cliente RPC (adicionar chamada a SomarIntervalo)
    solucoes/
      servidor.go            Solução do servidor (somente professor)
      cliente.go             Solução do cliente (somente professor)

Exercício 1 — Chat pub-sub com NATS:
  Servidor NATS: nats://demo.nats.io:4222 (público, sem autenticação)
  Testar conexão:  go run exercicio1-chat/teste-conexao.go
  Executar chat:   go run exercicio1-chat/chat.go <sala> <nome>
  Exemplo:
    Terminal 1:  go run exercicio1-chat/chat.go turma32 Alice
    Terminal 2:  go run exercicio1-chat/chat.go turma32 Bob

Exercício 2 — RPC com net/rpc:
  Iniciar servidor:  go run exercicio2-rpc/servidor.go
  Executar cliente:  go run exercicio2-rpc/cliente.go [endereco]

Requisitos:
  Go 1.22 ou superior — https://go.dev/dl
  Conexão com a internet (exercício 1: porta 4222 de demo.nats.io)
