# Go Examples

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Repositório com exemplos práticos organizados por tópico, cobrindo os principais conceitos da linguagem Go — da sintaxe básica a padrões mais avançados como generics e runtime.

## Estrutura do projeto

```
.
├── exampleArraysAndSlices/     
├── exampleErrors/              
├── exampleGoRunTime/           
├── exampleInterfaces/          
├── exampleMaps/                
├── exampleMethodAndStructs/    
├── exampleParametersGenerics/  
├── examplePointers/            
├── exampleReaderAndWriters/    
├── examplesArrays/             
├── examplesConditions/         
├── examplesDefer/              
├── examplesFunctions/          
├── examplesTypesAndConstants/  
├── examplesVariables/          
├── go.mod
└── main.go
```

## Tópicos abordados

| Pasta | Descrição |
|---|---|
| `exampleArraysAndSlices` | Diferenças entre arrays e slices, capacidade, append e manipulação |
| `exampleErrors` | Tratamento de erros, criação de erros customizados, `errors.Is`/`errors.As` |
| `exampleGoRunTime` | Conceitos de runtime do Go (goroutines, scheduler, garbage collector) |
| `exampleInterfaces` | Definição e implementação de interfaces, composição e uso idiomático |
| `exampleMaps` | Criação, manipulação e iteração de maps |
| `exampleMethodAndStructs` | Structs, métodos com receivers de valor e ponteiro |
| `exampleParametersGenerics` | Funções genéricas e uso de type parameters |
| `examplePointers` | Ponteiros, referências e manipulação de memória |
| `exampleReaderAndWriters` | Interfaces `io.Reader` e `io.Writer` e streams de dados |
| `examplesArrays` | Exemplos complementares sobre arrays |
| `examplesConditions` | Estruturas condicionais (`if`, `switch`) |
| `examplesDefer` | Uso de `defer`, `panic` e `recover` |
| `examplesFunctions` | Declaração de funções, retornos múltiplos, closures |
| `examplesTypesAndConstants` | Tipos customizados, `const`, `iota` |
| `examplesVariables` | Declaração e escopo de variáveis |

## Como executar

Cada pasta contém exemplos independentes. Para rodar um exemplo específico:

```bash
go run ./nomeDaPasta
```

Para rodar o arquivo principal:

```bash
go run main.go
```

## Requisitos

- Go instalado ([golang.org/dl](https://golang.org/dl))
- Verifique a versão utilizada no arquivo `go.mod`

## Objetivo

Este repositório serve como material de estudo e consulta rápida sobre os fundamentos e recursos da linguagem Go, útil tanto para revisão de conceitos quanto para preparação técnica.

---

*Sinta-se à vontade para adaptar os exemplos e adicionar novos tópicos conforme for aprofundando o estudo da linguagem.*
