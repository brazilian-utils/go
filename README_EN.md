
<div align="center">
<h1>🇧🇷 Brazilian Utils</h1>

<p>Utils library for Brazilian-specific businesses.</p>

[![CircleCI](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master.svg?style=svg)](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master)

### [Looking for the Portuguese version?](README.md)

</div>

# Introduction

Brazilian Utils is a library focused on solving problems we face daily in developing applications for Brazilian businesses.

- [Installation](#installation)
- [Usage](#usage)
- [Utilities](#utilities)
- [New Utilities and Reporting Bugs](#new-utilities-and-reporting-bugs)
- [Questions? Ideas?](#questions-ideas)
- [Contributing to the Project Code](#contributing-to-the-project-code)

# Installation

```shell
go get -u -v github.com/agaragon/brutils-go
```

# Usage

```golang
package main

import "github.com/agaragon/brutils-go/cpf"

func main() {
    if cpf.IsValid("40364478829") {
    }
}
```

# Utilities

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

Returns whether the provided CPF's verification digits match its base number. This function does not verify the existence of the CPF; it only validates the string format.

Arguments:
  * cpf (string): The CPF to be validated, an 11-digit string

Returns:
  * bool: True if the verification digits match the base number,
          False otherwise.

Example:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cpf"
)

func main() {
    if cpf.IsValid("82178537464") {
      fmt.Println("Valid CPF")
    }
    if cpf.IsValid("00011122233") {
      fmt.Println("Invalid CPF")
    }
}
```

## CNPJ

### IsValid

Checks whether the provided CNPJ's (Cadastro Nacional da Pessoa Jurídica) verification digits match its base number. The input must be a string of digits of the appropriate length. This function does not verify the existence of the CNPJ; it only validates the string format.

Arguments:
  * cnpj (string): The CNPJ to be validated.

Returns:
  * bool: True if the verification digits match the base number,
          False otherwise.

Example:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cnpj"
)

func main() {
    if cnpj.IsValid("03560714000142") {
      fmt.Println("Valid CNPJ")
    }

    if cnpj.IsValid("00111222000133") {
      fmt.Println("Invalid CNPJ")
    }
}
```

## CEP

### IsValid

Checks whether a Brazilian CEP (Postal Address Code) is valid. For a CEP to be considered valid, the input must be a string containing exactly 8 digits. This function does not verify if the CEP is a real CEP, as it only validates the string format.

Arguments:
  * cep (string): The string containing the CEP to be verified.

Returns:
  * bool: True if the CEP is valid (8 digits), False otherwise.

Example:

```golang
package main

import (
  "fmt"
  "github.com/agaragon/brutils-go/cep"
)

func main() {
    if cep.IsValid("00000-010") {
      fmt.Println("Valid CEP")
    }
    if cep.IsValid("00000-00000") {
      fmt.Println("Invalid CEP")
    }
}
```

### Clean

Removes any non-numeric characters from the CEP, returning only the numbers present.

Arguments:
  * cep (string): The string containing the CEP to be cleaned.

Returns:
  * string: The cleaned CEP string with only numeric characters.

Example:

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

Generates a random 8-digit CEP (Postal Address Code) as a string.

Returns:
  * string: A randomly generated 8-digit number with a separator dash.

Example:

```golang
package main

import (
    "fmt"
    "github.com/agaragon/brutils-go/cep"
)

func main() {
    fmt.Println(cep.Generate())
}
"12345-123"
```

### FetchAddress

Fetches the address corresponding to the provided CEP.

Returns:
  * Address: The address data corresponding to the provided CEP.

Example:

```golang
package main

import (
    "fmt"
    "github.com/agaragon/brutils-go/cep"
)

func main() {
    fmt.Println(cep.FetchAddress("01001-000"))
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

# New Utilities and Reporting Bugs

If you have any suggestions for new utilities or find any bugs, please open an issue on our [GitHub page](https://github.com/agaragon/brutils-go).

# Questions? Ideas?

Feel free to reach out if you have any questions or ideas. You can contact us through our [GitHub page](https://github.com/agaragon/brutils-go).

# Contributing to the Project Code

We welcome contributions! If you would like to contribute to the project, please fork the repository and submit a pull request. For major changes, please open an issue first to discuss what you would like to change.