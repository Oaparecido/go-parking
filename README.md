# 🚗 Go Parking Simulator

<p align="center">
  <img src="./img/gopher.png" alt="Go Mascot" width="150"/>
</p>

[![Go Version](https://img.shields.io/github/go-mod/go-version/Oaparecido/go-parking)](https://golang.org/)

Bem-vindo ao **Go Parking Manager**! 

🚀 Este projeto é um simulador de estacionamento criado em Go para praticar conceitos de programação orientada a objetos e gerenciamento de recursos.

---

## 📋 Sobre o projeto

Este simulador:
- Gerencia veículos dentro de um estacionamento.
- Simula entrada e saída de veículos.
- Gera na saída o valor final que deve ser cobrado.
- Informa a quantidade de Veículos estacionados e também a quantidade de estacionamentos registrados.

Tudo isso foi desenvolvido para praticar e aprender conceitos de **POO (Programação Orientada a Objetos)** no Go! 💡

---

## 🛠️ Funcionalidades

- 🚘 **Gerenciamento de veículos**:
  - Registro de veículos.
  - Registro de entrada e saída.
  - Consulta de veículos estacionados.

- 📏 **Gerenciamento do estacionamento**:
  - Registro de estacionamentos.
  - Configuração da capacidade máxima.
  - Consulta de veiculos estacionados.
  - Valor cobrado perante saída do veículo.

---

## 🧩 Estrutura do Projeto

```plaintext
.
├── main.go          # Ponto de entrada do programa
├── models/          # Estruturas e modelos do simulador
│   ├── car.go       # Estrutura e métodos para carros
│   ├── space.go     # Estrutura e métodos para vagas do estacionamento 
│   └── parking.go   # Lógica de negócios do estacionamento
├── interfaces/      # Contratos de negócio
│   └── vehicle.go   # Contrato de métodos que devem ser implementados para um veiculo
└── README.md        # Documentação do projeto
