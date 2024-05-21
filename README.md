<div align="center">
<h1>🇧🇷 Brazilian Utils</h1>

<p>Utils library for Brazilian-specific businesses.</p>

[![CircleCI](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master.svg?style=svg)](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master)

### [Looking for the english version?](README_EN.md)

</div>

# Introdução

Brazilian Utils é uma biblioteca com foco na resolução de problemas que enfrentamos diariamente no
desenvolvimento de aplicações para o business Brasileiro.

- [Instalação](#instalação)
- [Utilização](#utilização)
- [Utilitários](#utilitários)
- [Novos Utilitários e Reportar Bugs](#novos-utilitários-e-reportar-bugs)
- [Dúvidas? Ideias?](#dúvidas-ideias)
- [Contribuindo com o Código do Projeto](#contribuindo-com-o-código-do-projeto)


# Instalação

```shell
go get -u -v github.com/agaragon/brutils-go
```

# Utilização


```golang
package main

import "github.com/agaragon/brutils-go/cpf"

func main() {
    if cpf.IsValid("40364478829") {
    }
}
```

# Utilitários

- [CPF](#cpf)
  - [IsValid](#IsValid)
- [CNPJ](#cnpj)
  - [IsValid](#IsValid)
- [CEP](#cep)
  - [IsValid](#IsValid)
  - [Clean](#Clean)
  - [Generate](#Generate)
  - [FetchAddress](#FetchAddress)

## CPF

### IsValid

Retorna se os dígitos de verificação do CPF fornecido
correspondem ao seu número base. Esta função não verifica a existência do CPF;
ela apenas valida o formato da string.

Argumentos:
  * cpf (string): O CPF a ser validado, uma string de 11 dígitos

Retorna:
  * bool: Verdadeiro se os dígitos de verificação corresponderem ao número base,
          Falso caso contrário.

Exemplo:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cpf"
)

func main() {
    if cpf.IsValid("82178537464") {
      fmt.Println("CPF válido")
    }
    if cpf.IsValid("00011122233") {
      fmt.Println("CPF inválido")
    }
}
```

## CNPJ

### IsValid

Verifica se os dígitos de verificação do CNPJ (Cadastro Nacional da Pessoa
Jurídica) fornecido correspondem ao seu número base. A entrada deve ser uma
string de dígitos com o comprimento apropriado. Esta função não verifica a
existência do CNPJ; ela só valida o formato da string.

Argumentos:
  * cnpj (string): O CNPJ a ser validado.

Retorna:
  * bool: True se os dígitos de verificação corresponderem ao número base,
        False caso contrário.

Exemplo:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cnpj"
)

func main() {
    if cnpj.IsValid("03560714000142") {
      fmt.Println("CNPJ válido")
    }

    if cnpj.IsValid("00111222000133") {
      fmt.Println("CNPJ inválido")
    }
}
```

## CEP

### IsValid

Verifica se um CEP (Código de Endereçamento Postal) brasileiro é válido.
Para que um CEP seja considerado válido, a entrada deve ser uma string contendo
exatamente 8 dígitos. Esta função não verifica se o CEP é um CEP real, pois
valida apenas o formato da string.

Argumentos:
  * cep (string): A string contendo o CEP a ser verificado.

Retorno:
  * bool: True se o CEP for válido (8 dígitos), False caso contrário.

Exemplo:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cep"
)

func main() {
    if cep.IsValid("00000-010") {
      fmt.Println("CEP válido")
    }
    if cep.IsValid("00000-00000") {
      fmt.Println("CEP inválido")
    }
}
```

### Clean

Remove quaisquer caracteres não numéricos do CEP devolvendo apensa os números presentes.

Argumentos:
  * cep (string): A string contendo o CEP a ser verificado.

Retorno:
  * bool: True se o CEP for válido (8 dígitos), False caso contrário.

Exemplo:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cep"
)

func main() {
    fmt.Println(cep.Clean("00000-010"))
    fmt.Println(cep.Clean("00000000"))
}
"00000010"
"00000000"
```

### Generate

Gera um número de CEP (Código de Endereçamento Postal) aleatório de 8 dígitos
como uma string.

Retorna:
  * string: Um número de 8 dígitos gerado aleatoriamente com o traço separador.

Exemplo:

```golang
package main

import (
    "fmt"
    "github.com/agaragon/brutils-go/cpf"
)

func main() {
    fmt.Println(cpf.Generate())
}
"12345-123"
```

### FetchAddress

Gera busca o endereço correspondente ao CEP passado.

Retorna:
  * Address: Os dados do endereço correpondente ao cep passado.

Exemplo:

```golang
package main

import (
    "fmt"
    "github.com/agaragon/brutils-go/cpf"
)

func main() {
    fmt.Println(cpf.FetchAddress("01001-000"))
}

{
    "cep": "01001-000",
    "logradouro": "Praça da Sé",
    "complemento": "lado ímpar",
    "bairro": "Sé",
    "localidade": "São Paulo",
    "uf": "SP",
    "ibge": "3550308",
    "gia": "1004",
    "ddd": "11",
    "siafi": "7107"
}
```