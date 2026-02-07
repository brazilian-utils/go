# Brazilian Utils for Go / Utilitários Brasileiros para Go

[![CircleCI](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master.svg?style=svg)](https://circleci.com/gh/brazilian-utils/brutils-go/tree/master)
[![Go Version](https://img.shields.io/github/go-mod/go-version/brazilian-utils/brutils-go)](go.mod)

[🇺🇸 English](#english) | [🇧🇷 Português](#português)

---

## English

### 📖 About

**Brazilian Utils for Go** is a comprehensive library that provides utilities for working with Brazilian-specific data formats and documents. It includes validators, formatters, and generators for various Brazilian documents and data types.

### 🚀 Installation

```shell
go get -u github.com/brazilian-utils/brutils-go
```

### 📦 Packages

The library is organized into specialized packages, each handling a specific Brazilian data type:

- **[CPF](#cpf)** - Individual Taxpayer Registry
- **[CNPJ](#cnpj)** - National Registry of Legal Entities
- **[CEP](#cep)** - Postal Code
- **[Phone](#phone)** - Brazilian Phone Numbers
- **[Currency](#currency)** - Brazilian Real (R$)
- **[Boleto](#boleto)** - Payment Slip
- **[Email](#email)** - Email Address
- **[PIS](#pis)** - Social Integration Program
- **[CNH](#cnh)** - National Driver's License
- **[RENAVAM](#renavam)** - National Vehicle Registry
- **[License Plate](#license-plate)** - Vehicle License Plates
- **[Date](#date)** - Brazilian Date Utilities
- **[Legal Nature](#legal-nature)** - Legal Entity Nature
- **[Legal Process](#legal-process)** - Legal Process Numbers

---

### CPF

CPF (Cadastro de Pessoas Físicas) is the Brazilian individual taxpayer identification number.

```go
import "github.com/brazilian-utils/brutils-go/cpf"

// Validate CPF
cpf.IsValid("40364478829")  // true
cpf.IsValid("403.644.788-29")  // true
cpf.IsValid("00000000000")  // false (blacklisted)

// Format CPF
cpf.Format("40364478829")  // "403.644.788-29"
cpf.Format("403644788")  // "403.644.788" (incomplete)

// Generate random valid CPF
cpf.Generate()  // "12345678909" (random)
```

---

### CNPJ

CNPJ (Cadastro Nacional da Pessoa Jurídica) is the Brazilian company identification number.

```go
import "github.com/brazilian-utils/brutils-go/cnpj"

// Validate CNPJ
cnpj.IsValid("11222333000181")  // true
cnpj.IsValid("11.222.333/0001-81")  // true
cnpj.IsValid("00000000000000")  // false (blacklisted)

// Format CNPJ
cnpj.Format("11222333000181")  // "11.222.333/0001-81"

// Generate random valid CNPJ
cnpj.Generate()  // "12345678000190" (random)
```

---

### CEP

CEP (Código de Endereçamento Postal) is the Brazilian postal code.

```go
import "github.com/brazilian-utils/brutils-go/cep"

// Validate CEP
cep.IsValid("01310100")  // true
cep.IsValid("01310-100")  // false (must be digits only)

// Format CEP
cep.Format("01310100")  // "01310-100"
cep.Format("0131010")  // "" (invalid, returns empty)

// Generate random CEP
cep.Generate()  // "12345678" (random)

// Fetch address from CEP (uses ViaCEP API)
addr, err := cep.GetAddressFromCEP("01310100")
if err == nil {
    fmt.Println(addr.Logradouro)  // "Avenida Paulista"
    fmt.Println(addr.Bairro)      // "Bela Vista"
    fmt.Println(addr.Localidade)  // "São Paulo"
    fmt.Println(addr.UF)          // "SP"
}

// Search CEP from address
ceps, err := cep.GetCEPFromAddress("SP", "São Paulo", "Paulista")
if err == nil {
    for _, c := range ceps {
        fmt.Println(c.CEP, c.Logradouro)
    }
}
```

---

### Phone

Utilities for Brazilian phone numbers (mobile and landline).

```go
import "github.com/brazilian-utils/brutils-go/phone"

// Validate phone numbers
phone.IsValid("11987654321", "mobile")    // true
phone.IsValid("1133334444", "landline")   // true
phone.IsValid("11987654321", "")          // true (any type)

// Format phone numbers
phone.Format("11987654321")  // "(11)98765-4321"
phone.Format("1133334444")   // "(11)3333-4444"

// Remove symbols
phone.RemoveSymbols("(11) 98765-4321")  // "11987654321"

// Remove international code
phone.RemoveInternationalDialingCode("5511987654321")  // "11987654321"

// Generate random phone
phone.Generate("mobile")    // "11987654321" (random mobile)
phone.Generate("landline")  // "1133334444" (random landline)
phone.Generate("")          // random mobile or landline
```

---

### Currency

Format and convert Brazilian Real (R$) values.

```go
import "github.com/brazilian-utils/brutils-go/currency"

// Format currency
currency.FormatCurrency(1234.56)  // "R$ 1.234,56"
currency.FormatCurrency(1000000.00)  // "R$ 1.000.000,00"

// Convert to text (Portuguese)
currency.ConvertRealToText(1234.56)  // "Mil duzentos e trinta e quatro reais e cinquenta e seis centavos"
currency.ConvertRealToText(1.00)     // "Um real"
currency.ConvertRealToText(0.50)     // "Cinquenta centavos"
currency.ConvertRealToText(-100.00)  // "Menos cem reais"
```

---

### Boleto

Validate Brazilian bank payment slips (boletos).

```go
import "github.com/brazilian-utils/brutils-go/boleto"

// Validate boleto digitable line (47 digits)
boleto.IsValid("34191790010104351004791020150008291070026000")  // true/false
```

---

### Email

Email address validation.

```go
import "github.com/brazilian-utils/brutils-go/email"

// Validate email
email.IsValid("user@example.com")  // true
email.IsValid("invalid.email")     // false
email.IsValid(".user@example.com") // false (starts with dot)
```

---

### PIS

PIS (Programa de Integração Social) is a Brazilian social integration program number.

```go
import "github.com/brazilian-utils/brutils-go/pis"

// Validate PIS
pis.IsValid("12345678901")  // true/false

// Format PIS
pis.Format("12345678901")  // "123.45678.90-1"

// Generate random valid PIS
pis.Generate()  // "12345678901" (random)
```

---

### CNH

CNH (Carteira Nacional de Habilitação) is the Brazilian national driver's license.

```go
import "github.com/brazilian-utils/brutils-go/cnh"

// Validate CNH (11 digits)
cnh.IsValid("12345678901")  // true/false
cnh.IsValid("00000000000")  // false (blacklisted)
```

---

### RENAVAM

RENAVAM is the Brazilian national vehicle registration number.

```go
import "github.com/brazilian-utils/brutils-go/renavam"

// Validate RENAVAM (11 digits)
renavam.IsValid("12345678901")  // true/false
renavam.IsValid("11111111111")  // false (all same digit)
```

---

### License Plate

Utilities for Brazilian vehicle license plates (old format and Mercosul).

```go
import "github.com/brazilian-utils/brutils-go/licenseplate"

// Validate license plates
licenseplate.IsValid("ABC1234", "old_format")  // true
licenseplate.IsValid("ABC1D34", "mercosul")    // true
licenseplate.IsValid("ABC1D34", "")            // true (any format)

// Get format
licenseplate.GetFormat("ABC1234")  // "LLLNNNN"
licenseplate.GetFormat("ABC1D34")  // "LLLNLNN"

// Format license plate
licenseplate.Format("ABC1234")  // "ABC-1234" (old format with dash)
licenseplate.Format("abc1d34")  // "ABC1D34" (Mercosul uppercase)

// Convert old format to Mercosul
licenseplate.ConvertToMercosul("ABC1234")  // "ABC1B34"

// Generate random license plate
licenseplate.Generate("LLLNNNN")  // "ABC1234" (old format)
licenseplate.Generate("LLLNLNN")  // "ABC1D34" (Mercosul)
```

---

### Date

Convert Brazilian dates to Portuguese text.

```go
import "github.com/brazilian-utils/brutils-go/date"

// Convert date to text
date.ConvertDateToText("25/12/2024")  // "Vinte e cinco de Dezembro de dois mil e vinte e quatro"
date.ConvertDateToText("01/01/2000")  // "Primeiro de Janeiro de dois mil"
```

---

### Legal Nature

Legal Nature (Natureza Jurídica) codes from the Brazilian Federal Revenue.

```go
import "github.com/brazilian-utils/brutils-go/legalnature"

// Validate legal nature code
legalnature.IsValid("2062")  // true
legalnature.IsValid("206-2")  // true

// Get description
legalnature.GetDescription("2062")  // "Sociedade Empresária Limitada"
legalnature.GetDescription("2046")  // "Sociedade Anônima Aberta"
```

---

### Legal Process

Brazilian legal process number utilities.

```go
import "github.com/brazilian-utils/brutils-go/legalprocess"

// Validate legal process (20 digits)
legalprocess.IsValid("12345670820231234567")  // true/false

// Format legal process
legalprocess.Format("12345670820231234567")  // "1234567-08.2023.1.23.4567"

// Generate random legal process
legalprocess.Generate(2024, 1)  // "1234567082024123456" (random, year 2024, orgao 1)
```

---

### 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### 📄 License

This project is licensed under the MIT License.

---

## Português

### 📖 Sobre

**Brazilian Utils for Go** é uma biblioteca abrangente que fornece utilitários para trabalhar com formatos de dados e documentos específicos do Brasil. Inclui validadores, formatadores e geradores para vários documentos e tipos de dados brasileiros.

### 🚀 Instalação

```shell
go get -u github.com/brazilian-utils/brutils-go
```

### 📦 Pacotes

A biblioteca está organizada em pacotes especializados, cada um lidando com um tipo específico de dado brasileiro:

- **[CPF](#cpf-1)** - Cadastro de Pessoas Físicas
- **[CNPJ](#cnpj-1)** - Cadastro Nacional da Pessoa Jurídica
- **[CEP](#cep-1)** - Código de Endereçamento Postal
- **[Phone](#phone-1)** - Números de Telefone Brasileiros
- **[Currency](#currency-1)** - Real Brasileiro (R$)
- **[Boleto](#boleto-1)** - Boleto de Pagamento
- **[Email](#email-1)** - Endereço de Email
- **[PIS](#pis-1)** - Programa de Integração Social
- **[CNH](#cnh-1)** - Carteira Nacional de Habilitação
- **[RENAVAM](#renavam-1)** - Registro Nacional de Veículos Automotores
- **[License Plate](#license-plate-1)** - Placas de Veículos
- **[Date](#date-1)** - Utilitários de Data Brasileira
- **[Legal Nature](#legal-nature-1)** - Natureza Jurídica
- **[Legal Process](#legal-process-1)** - Números de Processos Judiciais

---

### CPF

CPF (Cadastro de Pessoas Físicas) é o número de identificação do contribuinte individual brasileiro.

```go
import "github.com/brazilian-utils/brutils-go/cpf"

// Validar CPF
cpf.IsValid("40364478829")  // true
cpf.IsValid("403.644.788-29")  // true
cpf.IsValid("00000000000")  // false (na lista negra)

// Formatar CPF
cpf.Format("40364478829")  // "403.644.788-29"
cpf.Format("403644788")  // "403.644.788" (incompleto)

// Gerar CPF válido aleatório
cpf.Generate()  // "12345678909" (aleatório)
```

---

### CNPJ

CNPJ (Cadastro Nacional da Pessoa Jurídica) é o número de identificação de empresa brasileiro.

```go
import "github.com/brazilian-utils/brutils-go/cnpj"

// Validar CNPJ
cnpj.IsValid("11222333000181")  // true
cnpj.IsValid("11.222.333/0001-81")  // true
cnpj.IsValid("00000000000000")  // false (na lista negra)

// Formatar CNPJ
cnpj.Format("11222333000181")  // "11.222.333/0001-81"

// Gerar CNPJ válido aleatório
cnpj.Generate()  // "12345678000190" (aleatório)
```

---

### CEP

CEP (Código de Endereçamento Postal) é o código postal brasileiro.

```go
import "github.com/brazilian-utils/brutils-go/cep"

// Validar CEP
cep.IsValid("01310100")  // true
cep.IsValid("01310-100")  // false (deve conter apenas dígitos)

// Formatar CEP
cep.Format("01310100")  // "01310-100"
cep.Format("0131010")  // "" (inválido, retorna vazio)

// Gerar CEP aleatório
cep.Generate()  // "12345678" (aleatório)

// Buscar endereço pelo CEP (usa a API ViaCEP)
addr, err := cep.GetAddressFromCEP("01310100")
if err == nil {
    fmt.Println(addr.Logradouro)  // "Avenida Paulista"
    fmt.Println(addr.Bairro)      // "Bela Vista"
    fmt.Println(addr.Localidade)  // "São Paulo"
    fmt.Println(addr.UF)          // "SP"
}

// Buscar CEP pelo endereço
ceps, err := cep.GetCEPFromAddress("SP", "São Paulo", "Paulista")
if err == nil {
    for _, c := range ceps {
        fmt.Println(c.CEP, c.Logradouro)
    }
}
```

---

### Phone

Utilitários para números de telefone brasileiros (celular e fixo).

```go
import "github.com/brazilian-utils/brutils-go/phone"

// Validar números de telefone
phone.IsValid("11987654321", "mobile")    // true
phone.IsValid("1133334444", "landline")   // true
phone.IsValid("11987654321", "")          // true (qualquer tipo)

// Formatar números de telefone
phone.Format("11987654321")  // "(11)98765-4321"
phone.Format("1133334444")   // "(11)3333-4444"

// Remover símbolos
phone.RemoveSymbols("(11) 98765-4321")  // "11987654321"

// Remover código internacional
phone.RemoveInternationalDialingCode("5511987654321")  // "11987654321"

// Gerar telefone aleatório
phone.Generate("mobile")    // "11987654321" (celular aleatório)
phone.Generate("landline")  // "1133334444" (fixo aleatório)
phone.Generate("")          // celular ou fixo aleatório
```

---

### Currency

Formatar e converter valores em Real Brasileiro (R$).

```go
import "github.com/brazilian-utils/brutils-go/currency"

// Formatar moeda
currency.FormatCurrency(1234.56)  // "R$ 1.234,56"
currency.FormatCurrency(1000000.00)  // "R$ 1.000.000,00"

// Converter para texto (português)
currency.ConvertRealToText(1234.56)  // "Mil duzentos e trinta e quatro reais e cinquenta e seis centavos"
currency.ConvertRealToText(1.00)     // "Um real"
currency.ConvertRealToText(0.50)     // "Cinquenta centavos"
currency.ConvertRealToText(-100.00)  // "Menos cem reais"
```

---

### Boleto

Validar boletos bancários brasileiros.

```go
import "github.com/brazilian-utils/brutils-go/boleto"

// Validar linha digitável do boleto (47 dígitos)
boleto.IsValid("34191790010104351004791020150008291070026000")  // true/false
```

---

### Email

Validação de endereço de email.

```go
import "github.com/brazilian-utils/brutils-go/email"

// Validar email
email.IsValid("user@example.com")  // true
email.IsValid("invalid.email")     // false
email.IsValid(".user@example.com") // false (começa com ponto)
```

---

### PIS

PIS (Programa de Integração Social) é um número do programa de integração social brasileiro.

```go
import "github.com/brazilian-utils/brutils-go/pis"

// Validar PIS
pis.IsValid("12345678901")  // true/false

// Formatar PIS
pis.Format("12345678901")  // "123.45678.90-1"

// Gerar PIS válido aleatório
pis.Generate()  // "12345678901" (aleatório)
```

---

### CNH

CNH (Carteira Nacional de Habilitação) é a carteira de motorista nacional brasileira.

```go
import "github.com/brazilian-utils/brutils-go/cnh"

// Validar CNH (11 dígitos)
cnh.IsValid("12345678901")  // true/false
cnh.IsValid("00000000000")  // false (na lista negra)
```

---

### RENAVAM

RENAVAM é o número do registro nacional de veículos automotores brasileiro.

```go
import "github.com/brazilian-utils/brutils-go/renavam"

// Validar RENAVAM (11 dígitos)
renavam.IsValid("12345678901")  // true/false
renavam.IsValid("11111111111")  // false (todos dígitos iguais)
```

---

### License Plate

Utilitários para placas de veículos brasileiros (formato antigo e Mercosul).

```go
import "github.com/brazilian-utils/brutils-go/licenseplate"

// Validar placas
licenseplate.IsValid("ABC1234", "old_format")  // true
licenseplate.IsValid("ABC1D34", "mercosul")    // true
licenseplate.IsValid("ABC1D34", "")            // true (qualquer formato)

// Obter formato
licenseplate.GetFormat("ABC1234")  // "LLLNNNN"
licenseplate.GetFormat("ABC1D34")  // "LLLNLNN"

// Formatar placa
licenseplate.Format("ABC1234")  // "ABC-1234" (formato antigo com traço)
licenseplate.Format("abc1d34")  // "ABC1D34" (Mercosul maiúsculo)

// Converter formato antigo para Mercosul
licenseplate.ConvertToMercosul("ABC1234")  // "ABC1B34"

// Gerar placa aleatória
licenseplate.Generate("LLLNNNN")  // "ABC1234" (formato antigo)
licenseplate.Generate("LLLNLNN")  // "ABC1D34" (Mercosul)
```

---

### Date

Converter datas brasileiras para texto em português.

```go
import "github.com/brazilian-utils/brutils-go/date"

// Converter data para texto
date.ConvertDateToText("25/12/2024")  // "Vinte e cinco de Dezembro de dois mil e vinte e quatro"
date.ConvertDateToText("01/01/2000")  // "Primeiro de Janeiro de dois mil"
```

---

### Legal Nature

Códigos de Natureza Jurídica da Receita Federal Brasileira.

```go
import "github.com/brazilian-utils/brutils-go/legalnature"

// Validar código de natureza jurídica
legalnature.IsValid("2062")  // true
legalnature.IsValid("206-2")  // true

// Obter descrição
legalnature.GetDescription("2062")  // "Sociedade Empresária Limitada"
legalnature.GetDescription("2046")  // "Sociedade Anônima Aberta"
```

---

### Legal Process

Utilitários para números de processos judiciais brasileiros.

```go
import "github.com/brazilian-utils/brutils-go/legalprocess"

// Validar processo judicial (20 dígitos)
legalprocess.IsValid("12345670820231234567")  // true/false

// Formatar processo judicial
legalprocess.Format("12345670820231234567")  // "1234567-08.2023.1.23.4567"

// Gerar processo judicial aleatório
legalprocess.Generate(2024, 1)  // "1234567082024123456" (aleatório, ano 2024, órgão 1)
```

---

### 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para enviar um Pull Request.

### 📄 Licença

Este projeto está licenciado sob a Licença MIT.
