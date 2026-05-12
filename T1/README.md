# T1 — Jogo Concorrente de terminal em GO

**Disciplina:** Fundamentos de Processamento Paralelo e Distribuído (FPPD)

## 👥 Integrantes do Grupo

* Emile Vargas Bordin
* Felipe Souza Flores
* Natan de Aragão Müller

---

## 🎯 Objetivo

Este projeto tem como objetivo aplicar os conceitos de **programação concorrente** estudados ao longo da disciplina, desenvolvendo um jogo interativo em terminal utilizando a linguagem **Go**, onde a concorrência não é apenas um detalhe, mas parte essencial da arquitetura.

---

## 🎮 Jogo Escolhido

O projeto desenvolvido se trata de uma adaptação do jogo **"Taco Gato Cabra Queijo Pizza"**, um jogo de reflexo e atenção onde:

* Os jogadores revelam cartas em sequência

* A cada jogada, uma palavra da sequência é dita:

  **Taco → Gato → Cabra → Queijo → Pizza → (repete)**

* Quando a carta revelada coincide com a palavra falada:

  * Todos os jogadores devem reagir rapidamente (bater no monte)
  * O último a reagir sofre penalidade

A proposta do projeto é simular esse comportamento em ambiente concorrente, com múltiplos jogadores executando ações de forma independente.

---

## 🧠 Conceitos Aplicados

O projeto utiliza conceitos fundamentais de concorrência, como:

* Goroutines (execução concorrente)
* Channels (comunicação entre processos)
* Select (multiplexação de eventos)
* Sincronização e coordenação entre processos
* Condição de parada (shutdown gracioso)

---

## 📂 Estrutura do Projeto

```text
T1/
│
├── Arquitetura/
│   └── DocDeArquitetura.pdf
│
├── Enunciado/
│   └── ENUNCIADO.pdf
│
├── Implementação/
│   ├── main.go
│   ├── game.go
│   ├── player.go
│   └── deck.go
│
├── Material teórico/
│   ├── 1 - Introdução.pdf
│   ├── 2 - Arquiteturas MIMD.pdf
│   ├── 3 - Conceitos fundamentais de sistemas concorrentes.pdf
│   ├── 4 - Programação concorrente em GO.pdf
│   ├── 5 - Comunicação e sincronização em sistemas concorrentes.pdf
│   ├── 6 - Canais e memória compartilhada.pdf
│   └── 7 - Mecanismos de sincronização.pdf
│
└── README.md
```

### 📌 Organização dos Arquivos

* **main.go**
  Responsável por inicializar o sistema, criar as goroutines e coordenar o fluxo principal da aplicação. Define quatro goroutines concorrentes: `gameLoop()` (controla a lógica do jogo), `renderer()` (exibe eventos em tempo real), `inputPlayer()` (captura entrada do usuário) e `automaticTurns()` (executa turnos automáticos). Utiliza channels para comunicação entre as goroutines e um canal `done` para shutdown gracioso.

* **game.go**
  Contém a lógica central do jogo, incluindo controle de turnos, validação de regras e fluxo da partida.

* **player.go**
  Implementa o comportamento dos jogadores, incluindo ações e reações.

* **deck.go**
  Responsável pela criação, organização e manipulação do baralho/cartas do jogo.

---

## ⚙️ Requisitos do Trabalho

O projeto atende aos seguintes requisitos definidos no enunciado:

* ✔ Pelo menos **4 goroutines com papéis distintos**
* ✔ Comunicação entre goroutines feita por **channels**
* ✔ Uso de **select** para multiplexação de eventos
* ✔ Presença de **entidades autônomas concorrentes**
* ✔ Implementação de **shutdown gracioso**
* ✔ Execução com **visualização em tempo real no terminal**
* ✔ Execução sem erros com:

  ```bash
  go run -race .
  ```

---

## 🕹️ Controles e Comandos

Durante o jogo, os seguintes comandos estão disponíveis:

* **ENTER** - Revelar carta do jogador humano (ou esperar 2 segundos pela revelação automática)
* **s** - Bater no monte (quando houver match)
* **q** - Sair do jogo

**Nota:** O jogador humano pode revelar sua carta manualmente pressionando ENTER ou aguardar a revelação automática que ocorre a cada 2 segundos.

---

## 🏗️ Implementação

Inicialmente, adotamos a abordagem de considerar um baralho de 36 cartas distribuídas igualmente em um número fixo de 3 jogadores (1 player + 2 bots), totalizando 12 para cada um.

**main.go**
  Responsável por inicializar o sistema, criar os channels e coordenar o fluxo principal da aplicação. Define as goroutines principais do sistema, incluindo `gameLoop()` (controla a lógica do jogo), `renderer()` (exibe eventos em tempo real), `inputPlayer()` (captura entrada do usuário), `automaticTurns()` (executa turnos automáticos) e as goroutines dos bots com `botPlayer()`. Utiliza channels para comunicação entre as goroutines e um canal `done` para shutdown gracioso.

O arquivo `deck.go` é responsável por definir o tipo `Card`, declarar as cartas do jogo e criar o baralho. A função `newDeck()` monta um slice de 36 cartas, embaralha os elementos usando `rand.Shuffle` e entrega um deck pronto para ser distribuído.

O arquivo `player.go` define a estrutura de jogadores e comandos do jogo. Ele também cria os três participantes iniciais (1 jogador humano e 2 bots) e implementa o comportamento de bots através da função `botPlayer()` que aguarda um tempo randômico antes de enviar ações ao canal de comando.

O arquivo `game.go` contém a lógica central do jogo, incluindo as estruturas `GameEvent` e `GameState` para gerenciar mensagens e o estado da partida. A função `gameLoop()` coordena o fluxo da partida usando `select` para multiplexar comandos e sinais de encerramento, sendo o único ponto que modifica o estado do jogo (garantindo consistência). A função `playTurn()` revela cartas, compara com a sequência e avança o turno. A função `handleSlap()` gerencia as batidas no monte, validando timing e determinando o perdedor da rodada.

---

## 📊 Documento de Arquitetura

O Documento detalhado de Arquitetura pode ser encontrado em:

```
/Arquitetura/DocDeArquitetura.pdf
```


## 📌 Observações

Este projeto foi desenvolvido priorizando o aprendizado dos conceitos de concorrência, seguindo os materiais próprios da disciplina.

