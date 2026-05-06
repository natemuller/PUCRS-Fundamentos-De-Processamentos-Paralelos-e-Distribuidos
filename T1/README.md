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
├── Diagrama/
│   └── Diagrama.pdf
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
  Responsável por inicializar o sistema, criar as goroutines e coordenar o fluxo principal da aplicação.

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
  go run -race
  ```

---

## 🏗️ Implementação

...

---

## 📊 Diagrama visual

O diagrama de goroutines e channels, pode ser encontrado em:

```
/Diagrama/Diagrama.pdf
```


## 📌 Observações

Este projeto foi desenvolvido priorizando o aprendizado dos conceitos de concorrência, seguindo os materiais próprios da disciplina.

