# Brazilian Utils for Go / Utilitários Brasileiros para Go

[![Go Version](https://img.shields.io/github/go-mod/go-version/brazilian-utils/brutils-go)](go.mod)

[🇧🇷 Português](#português) | [🇺🇸 English](#english)

---

## Português

### 📖 Sobre

**Brazilian Utils for Go** é uma biblioteca abrangente que fornece utilitários para trabalhar com formatos de dados e documentos específicos do Brasil. Inclui validadores, formatadores e geradores para vários documentos e tipos de dados brasileiros.

### 🚀 Instalação

```shell
go get -u github.com/brazilian-utils/go
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
- **[Voter ID](#voter-id-1)** - Título de Eleitor

**Documentos pessoais e profissionais:**
- **[CNS](#cns-1)** - Cartão Nacional de Saúde
- **[Certidão](#certidão-1)** - Certidão de Registro Civil
- **[Passport](#passport-1)** - Passaporte Brasileiro
- **[Professional Registration](#professional-registration-1)** - Registro em Conselhos Profissionais
- **[CBO](#cbo-1)** - Classificação Brasileira de Ocupações

**Cadastros de obra/empregador:**
- **[CEI](#cei-1)** - Cadastro Específico do INSS
- **[CNO](#cno-1)** - Cadastro Nacional de Obras

**Cadastros e códigos fiscais:**
- **[CAEPF](#caepf-1)** - Cadastro de Atividade Econômica da Pessoa Física
- **[IE](#ie-1)** - Inscrição Estadual
- **[CFOP](#cfop-1)** - Código Fiscal de Operações e Prestações
- **[CNAE](#cnae-1)** - Classificação Nacional de Atividades Econômicas
- **[CSOSN](#csosn-1)** - Código de Situação da Operação no Simples Nacional
- **[CST](#cst-1)** - Código de Situação Tributária
- **[NCM](#ncm-1)** - Nomenclatura Comum do Mercosul
- **[NFe Key](#nfe-key-1)** - Chave de Acesso de Documentos Fiscais Eletrônicos

**Financeiro e bancário:**
- **[Bank](#bank-1)** - Bancos Brasileiros (COMPE)
- **[Bank Account](#bank-account-1)** - Conta Bancária Brasileira
- **[IBAN](#iban-1)** - IBAN Brasileiro
- **[Credit Card](#credit-card-1)** - Validação de Cartão de Pagamento
- **[Pix Key](#pix-key-1)** - Chave Pix
- **[Pix Payload](#pix-payload-1)** - Payload do BR Code Pix

**Localização:**
- **[Area Code](#area-code-1)** - DDDs Brasileiros
- **[Municipality](#municipality-1)** - Municípios Brasileiros (IBGE)
- **[State](#state-1)** - Unidades Federativas Brasileiras

**Veículos:**
- **[VIN](#vin-1)** - Chassi de Veículo (Vehicle Identification Number)

**Texto e números:**
- **[Text](#text-1)** - Utilitários de Texto em Português
- **[Number](#number-1)** - Número por Extenso

---

### CPF

CPF (Cadastro de Pessoas Físicas) é o número de identificação do contribuinte individual brasileiro.

```go
import "github.com/brazilian-utils/go/cpf"

// Validar CPF
cpf.IsValid("40364478829")  // true
cpf.IsValid("403.644.788-29")  // true
cpf.IsValid("00000000000")  // false (na lista negra)

// Formatar CPF
cpf.Format("40364478829")  // "403.644.788-29"
cpf.Format("403644788")  // "403.644.788" (incompleto)

// Gerar CPF válido aleatório
cpf.Generate()  // "12345678909" (aleatório)

// Remover símbolos de formatação
cpf.RemoveSymbols("831.595.621-31")  // "83159562131"

// Extrair somente os dígitos (normalizar, truncando em 11)
cpf.Parse("943.895.751-04")  // "94389575104"
```

---

### CNPJ

CNPJ (Cadastro Nacional da Pessoa Jurídica) é o número de identificação de empresa brasileiro.

```go
import "github.com/brazilian-utils/go/cnpj"

// Validar CNPJ
cnpj.IsValid("11222333000181")  // true
cnpj.IsValid("11.222.333/0001-81")  // true
cnpj.IsValid("00000000000000")  // false (na lista negra)

// Formatar CNPJ
cnpj.Format("11222333000181")  // "11.222.333/0001-81"

// Gerar CNPJ válido aleatório
cnpj.Generate()  // "12345678000190" (aleatório)

// Gerar e validar um CNPJ alfanumérico (v2)
alnum := cnpj.Generate(cnpj.GenerateOptions{Version: 2})
cnpj.IsValid(alnum, cnpj.IsValidOptions{Version: 2})  // true

// Remover símbolos de formatação
cnpj.RemoveSymbols("10.799.163/9892-71")  // "10799163989271"

// Extrair somente os dígitos (normalizar, truncando em 14)
cnpj.Parse("46.843.485/0001-86")  // "46843485000186"
```

---

### CEP

CEP (Código de Endereçamento Postal) é o código postal brasileiro.

```go
import "github.com/brazilian-utils/go/cep"

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

// Remover símbolos de formatação
cep.RemoveSymbols("91906-292")  // "91906292"

// Extrair somente os dígitos (normalizar, truncando em 8)
cep.Parse("01001-000")  // "01001000"
```

---

### Phone

Utilitários para números de telefone brasileiros (celular e fixo).

```go
import "github.com/brazilian-utils/go/phone"

// Validar números de telefone (por padrão aceita celular e fixo)
phone.IsValid("11994029275")  // true
phone.IsValid("1635014415")   // true

// Restringir o tipo aceito
phone.IsValid("11994029275", phone.IsValidOptions{Kinds: []string{"mobile"}})    // true
phone.IsValid("11994029275", phone.IsValidOptions{Kinds: []string{"landline"}})  // false

// Validar um tipo específico diretamente
phone.IsValidMobile("(11) 98765-4321")     // true
phone.IsValidLandline("(11) 3000-0000")    // true
phone.IsValidService("0800 123 4567")      // true (número de utilidade pública/0800)

// Formatar números de telefone
phone.Format("988887777")                          // "98888-7777" (somente o número, sem DDD)
phone.Format("1130000000", phone.FormatOptions{Mask: "ddd"})  // "(11)3000-0000"

// Remover símbolos
phone.RemoveSymbols("(11) 99402-9275")  // "11994029275"

// Remover código internacional
phone.RemoveInternationalDialingCode("5511994029275")  // "11994029275"

// Extrair somente os dígitos (normalizar, com DDI removido quando aplicável)
phone.Parse("+55 11 98888-7777")  // "11988887777"

// Gerar telefone aleatório
phone.Generate("mobile")    // celular aleatório
phone.Generate("landline")  // fixo aleatório
phone.Generate("")          // celular ou fixo aleatório
```

---

### Currency

Formatar e converter valores em Real Brasileiro (R$).

```go
import "github.com/brazilian-utils/go/currency"

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
import "github.com/brazilian-utils/go/boleto"

// Validar linha digitável do boleto (47 dígitos)
boleto.IsValid("34191790010104351004791020150008291070026000")  // true/false

// Formatar a linha digitável
boleto.Format("00190000090114971860168524522114675860000102656")
// "00190.00009 01149.718601 68524.522114 6 75860000102656"

// Extrair somente os dígitos (normalizar)
boleto.Parse("00190.00009 01149.718601 68524.522114 6 75860000102656")
// "00190000090114971860168524522114675860000102656"

// Gerar um boleto (cobrança bancária) válido aleatório
boleto.Generate()

// Extrair banco, valor e data de vencimento
info := boleto.GetInfo("00190000090114971860168524522114675860000102656")
fmt.Println(info.BankCode, info.AmountCents, info.DueDate)
// "001" 102656 2018-07-15
```

---

### Email

Validação de endereço de email.

```go
import "github.com/brazilian-utils/go/email"

// Validar email
email.IsValid("user@example.com")  // true
email.IsValid("invalid.email")     // false
email.IsValid(".user@example.com") // false (começa com ponto)
```

---

### PIS

PIS (Programa de Integração Social) é um número do programa de integração social brasileiro.

```go
import "github.com/brazilian-utils/go/pis"

// Validar PIS
pis.IsValid("12345678901")  // true/false

// Formatar PIS
pis.Format("12345678901")  // "123.45678.90-1"

// Gerar PIS válido aleatório
pis.Generate()  // "12345678901" (aleatório)

// Remover símbolos de formatação
pis.RemoveSymbols("559.84775.36-3")  // "55984775363"

// Extrair somente os dígitos (normalizar, truncando em 11)
pis.Parse("123.45678.90-1")  // "12345678901"
```

---

### CNH

CNH (Carteira Nacional de Habilitação) é a carteira de motorista nacional brasileira.

```go
import "github.com/brazilian-utils/go/cnh"

// Validar CNH (11 dígitos)
cnh.IsValid("12345678901")  // true/false
cnh.IsValid("00000000000")  // false (na lista negra)
```

---

### RENAVAM

RENAVAM é o número do registro nacional de veículos automotores brasileiro.

```go
import "github.com/brazilian-utils/go/renavam"

// Validar RENAVAM (11 dígitos)
renavam.IsValid("12345678901")  // true/false
renavam.IsValid("11111111111")  // false (todos dígitos iguais)
```

---

### License Plate

Utilitários para placas de veículos brasileiros (formato antigo e Mercosul).

```go
import "github.com/brazilian-utils/go/licenseplate"

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

// Remover símbolos de formatação
licenseplate.RemoveSymbols("WDG1S21")  // "WDG1S21"

// Extrair e normalizar (maiúsculas, truncando em 7 caracteres)
licenseplate.Parse("abc-1234")  // "ABC1234"

// Gerar placa aleatória
licenseplate.Generate("LLLNNNN")  // "ABC1234" (formato antigo)
licenseplate.Generate("LLLNLNN")  // "ABC1D34" (Mercosul)
```

---

### Date

Converter datas brasileiras para texto em português.

```go
import "github.com/brazilian-utils/go/date"

// Converter data para texto
date.ConvertDateToText("25/12/2024")  // "Vinte e cinco de Dezembro de dois mil e vinte e quatro"
date.ConvertDateToText("01/01/2000")  // "Primeiro de Janeiro de dois mil"

// Verificar se uma data é feriado (nacional e, opcionalmente, estadual)
date.IsHoliday(date.IsHolidayOptions{Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)})  // true

// Listar os feriados nacionais de um ano
date.GetHolidays(2024)

// Verificar se é dia útil e somar/subtrair dias úteis
date.IsBusinessDay(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC))  // true (terça-feira)
date.AddBusinessDays(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 1)
// aponta para 2024-01-02 (pula o feriado de Ano Novo)

// Diferença em dias úteis entre duas datas
date.DifferenceInBusinessDays(
    time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
)  // 4
```

---

### Legal Nature

Códigos de Natureza Jurídica da Receita Federal Brasileira.

```go
import "github.com/brazilian-utils/go/legalnature"

// Validar código de natureza jurídica
legalnature.IsValid("2062")  // true
legalnature.IsValid("206-2")  // true

// Obter descrição
legalnature.GetDescription("2062")  // "Sociedade Empresária Limitada"
legalnature.GetDescription("2046")  // "Sociedade Anônima Aberta"

// Consultar a entrada completa (código, descrição e categoria)
ln := legalnature.Get("2062")
fmt.Println(ln.Description, ln.Category.Code)  // "Sociedade Empresária Limitada" "2"

// Listar todos os códigos em vigor (92) e por categoria CONCLA
legalnature.List()
legalnature.ListByCategory("5")

// Formatar, normalizar e gerar um código aleatório válido
legalnature.Format("2062")  // "206-2"
legalnature.Parse("206-2")  // "2062"
legalnature.Generate()      // código aleatório dentre os vigentes
```

---

### Legal Process

Utilitários para números de processos judiciais brasileiros.

```go
import "github.com/brazilian-utils/go/legalprocess"

// Validar processo judicial (20 dígitos)
legalprocess.IsValid("12345670820231234567")  // true/false

// Formatar processo judicial
legalprocess.Format("12345670820231234567")  // "1234567-08.2023.1.23.4567"

// Gerar processo judicial aleatório
legalprocess.Generate(legalprocess.GenerateOptions{Court: 1})  // aleatório, ano atual, órgão 1

// Remover símbolos de formatação
legalprocess.RemoveSymbols("1666923-98.2026.9.26.4931")  // "16669239820269264931"

// Extrair somente os dígitos (normalizar, truncando em 20)
legalprocess.Parse("0002080-25.2012.5.15.0049")  // "00020802520125150049"
```

---

### Voter ID

Título de Eleitor (Voter ID) é o número de identificação do eleitor brasileiro.

```go
import "github.com/brazilian-utils/go/voterid"

// Validar Título de Eleitor
voterid.IsValid("217633460930")  // true
voterid.IsValid("6908 4709 28 28")  // true (com espaços)
voterid.IsValid("123456789011")  // false (dígitos verificadores inválidos)

// Formatar Título de Eleitor
voterid.Format("217633460930")  // "2176 3346 09 30"
voterid.Format("3244567800167")  // "3244 5678 00 16" (13 dígitos para SP/MG)

// Gerar Título de Eleitor válido aleatório
voterid.Generate()       // "217633460930" (aleatório, padrão ZZ - estrangeiro)
voterid.Generate("SP")   // "149426030183" (São Paulo)
voterid.Generate("MG")   // "033568860230" (Minas Gerais)
voterid.Generate("RJ")   // "858224120973" (Rio de Janeiro)
```

**Observações:**
- Títulos de eleitor têm normalmente 12 dígitos
- Casos especiais: SP (01) e MG (02) podem ter 13 dígitos
- Códigos de UF válidos: 01-28
- Dois dígitos verificadores calculados usando módulo 11

---

### CNS

CNS (Cartão Nacional de Saúde) é o número do cartão do SUS do cidadão brasileiro.

```go
import "github.com/brazilian-utils/go/cns"

// Validar CNS
cns.IsValid("123456789010000")     // true
cns.IsValid("123 4567 8901 0000")  // true (com espaços)
cns.IsValid("700000000000005")     // true (cartão provisório)

// Formatar CNS
cns.Format("123456789010001")  // "123 4567 8901 0001"

// Remover formatação
cns.Parse("898 0000 0004 3208")  // "898000000043208"
```

---

### Certidão

Validação, formatação e leitura da matrícula de certidão de registro civil (nascimento, casamento, óbito etc.), conforme o art. 473 do Código Nacional de Normas da Corregedoria Nacional de Justiça.

```go
import "github.com/brazilian-utils/go/certidao"

// Validar matrícula (32 dígitos)
certidao.IsValid("10453901552013100012021000012321")  // true

// Formatar matrícula
certidao.Format("10453901552013100012021000012321")
// "104539 01 55 2013 1 00012 021 0000123 21"

// Extrair informações da matrícula
info := certidao.GetInfo("104539 01 55 2013 1 00012 021 0000123 21")
fmt.Println(info.Year, info.Type)  // 2013 "birth"
```

---

### Passport

Utilitários para o número do passaporte brasileiro (2 letras seguidas de 6 dígitos).

```go
import "github.com/brazilian-utils/go/passport"

// Validar passaporte
passport.IsValid("AA111111")   // true
passport.IsValid("AB-123456")  // true (com símbolos)

// Formatar/normalizar
passport.Format("acd12736")  // "ACD12736"

// Gerar passaporte aleatório
passport.Generate()  // "AB123456" (aleatório)
```

---

### Professional Registration

Verificação estrutural (quantidade de dígitos e UF) do registro em conselhos profissionais: OAB, CRM, CRO, CRP e CRC. Não há dígito verificador publicado para nenhum deles.

```go
import "github.com/brazilian-utils/go/registroprofissional"

// Validar registro na OAB
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "123456/SP", Council: "OAB",
})  // true

// Restringir por estado
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "123456/SP", Council: "OAB", State: "RJ",
})  // false

// Validar registro no CRC (com o dígito de tipo O/P)
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "SP-123456/O-3", Council: "CRC",
})  // true
```

---

### CBO

CBO (Classificação Brasileira de Ocupações) é a tabela oficial de códigos de ocupação (CBO 2002).

```go
import "github.com/brazilian-utils/go/cbo"

// Validar código CBO
cbo.IsValid("212405")   // true
cbo.IsValid("2124-05")  // true (com máscara)

// Consultar título oficial
c := cbo.Get("212405")
fmt.Println(c.Description)  // "Analista de desenvolvimento de sistemas"

// Remover formatação
cbo.Parse("2124-05")  // "212405"
```

---

### CEI

CEI (Cadastro Específico do INSS) é o número de cadastro de obras/empregadores, hoje substituído pelo CNO (mas ainda em uso).

```go
import "github.com/brazilian-utils/go/cei"

// Validar CEI
cei.IsValid("11.583.00249/85")  // true
cei.IsValid("115830024985")     // true

// Formatar CEI
cei.Format("277297118187")  // "27.729.71181/87"

// Remover formatação
cei.Parse("27.729.71181/87")  // "277297118187"
```

---

### CNO

CNO (Cadastro Nacional de Obras) é o número de cadastro de obras que substituiu o CEI, mantendo a mesma regra de dígitos.

```go
import "github.com/brazilian-utils/go/cno"

// Validar CNO
cno.IsValid("110840168062")     // true
cno.IsValid("11.084.01680/62")  // true

// Formatar CNO
cno.Format("111130137368")  // "11.113.01373/68"
```

---

### CAEPF

CAEPF (Cadastro de Atividade Econômica da Pessoa Física) é o número de cadastro de atividade econômica da pessoa física.

```go
import "github.com/brazilian-utils/go/caepf"

// Validar CAEPF
caepf.IsValid("293.118.610/001-84")  // true
caepf.IsValid("29311861000184")      // true

// Formatar CAEPF
caepf.Format("29311861000184")  // "293.118.610/001-84"

// Remover formatação
caepf.Parse("293.118.610/001-84")  // "29311861000184"
```

---

### IE

IE (Inscrição Estadual) é o número de inscrição estadual do contribuinte. A validação é estrutural (quantidade de dígitos por UF), sem dígito verificador (ver documentação do pacote).

```go
import "github.com/brazilian-utils/go/ie"

// Validar IE por estado
ie.IsValid(ie.IsValidParams{Value: "123456789012", State: "SP"})   // true
ie.IsValid(ie.IsValidParams{Value: "12345678", State: "RJ"})       // true
ie.IsValid(ie.IsValidParams{Value: "P123456789012", State: "SP"})  // true (produtor rural)
```

---

### CFOP

CFOP (Código Fiscal de Operações e Prestações) é o código fiscal oficial das operações e prestações (Convênio SINIEF s/nº 1970, Anexo II consolidado).

```go
import "github.com/brazilian-utils/go/cfop"

// Validar código CFOP
cfop.IsValid("5102")   // true
cfop.IsValid("5.102")  // true (com máscara)

// Consultar descrição oficial
c := cfop.Get("5102")
fmt.Println(c.Description)

// Remover formatação
cfop.Parse("5.102")  // "5102"
```

---

### CNAE

CNAE (Classificação Nacional de Atividades Econômicas) é o código oficial de subclasse de atividade econômica (CNAE-Subclasses 2.3).

```go
import "github.com/brazilian-utils/go/cnae"

// Validar código CNAE
cnae.IsValid("6201501")    // true
cnae.IsValid("6201-5/01")  // true (com máscara)

// Formatar código CNAE
cnae.Format("6201501")  // "6201-5/01"

// Consultar descrição oficial
c := cnae.Get("6201501")
fmt.Println(c.Description)  // "DESENVOLVIMENTO DE PROGRAMAS DE COMPUTADOR SOB ENCOMENDA"
```

---

### CSOSN

CSOSN (Código de Situação da Operação no Simples Nacional) é o código oficial das 10 situações do Simples Nacional (Ajuste SINIEF 07/2005).

```go
import "github.com/brazilian-utils/go/csosn"

// Validar código CSOSN
csosn.IsValid("101")  // true
csosn.IsValid(101)    // true (número)
csosn.IsValid("999")  // false (fora da tabela)
```

---

### CST

CST (Código de Situação Tributária) é o código oficial de situação tributária do ICMS, IPI e PIS/COFINS.

```go
import "github.com/brazilian-utils/go/cst"

// Validar CST (qualquer tabela)
cst.IsValid("110")  // true (ICMS)
cst.IsValid("07")   // true (PIS/COFINS)

// Restringir a uma tabela específica
cst.IsValid("07", cst.IsValidOptions{Tax: "ipi"})  // false
```

---

### NCM

NCM (Nomenclatura Comum do Mercosul) é o código oficial de classificação de mercadorias da tabela vigente do Siscomex.

```go
import "github.com/brazilian-utils/go/ncm"

// Validar código NCM
ncm.IsValid("22030000")    // true
ncm.IsValid("2203.00.00")  // true (com máscara)

// Formatar código NCM
ncm.Format("84713012")  // "8471.30.12"
```

---

### NFe Key

Chave de acesso de documentos fiscais eletrônicos: NF-e, NFC-e, CT-e, MDF-e, CT-e OS, GTV-e, BP-e, NF3e e NFCom.

```go
import "github.com/brazilian-utils/go/nfekey"

key := "35170458716523000119550010000000121000123458"

// Validar chave de acesso
nfekey.IsValid(key)  // true

// Formatar chave de acesso
nfekey.Format(key)
// "3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458"

// Extrair informações da chave
info := nfekey.GetInfo(key)
fmt.Println(info.StateCode, info.Year, info.Model)  // "SP" 2017 "55"
```

---

### Bank

Consulta aos bancos brasileiros com código COMPE, publicados pelo Banco Central na lista de participantes do STR.

```go
import "github.com/brazilian-utils/go/bank"

// Buscar banco pelo código COMPE
b := bank.GetByCode("341")
fmt.Println(b.Name)  // "ITAÚ UNIBANCO S.A."

// Buscar banco pelo ISPB
bank.GetByIspb("00000000")  // Banco do Brasil (código 001)

// Listar todos os bancos com código COMPE
bank.List()
```

---

### Bank Account

Validação de conta bancária brasileira: código do banco, agência, conta e dígito verificador. Implementação parcial (ver documentação do pacote para o que é coberto).

```go
import "github.com/brazilian-utils/go/bankaccount"

// Validar conta bancária
bankaccount.IsValid(bankaccount.IsValidParams{
    BankCode: "341", Agency: "1234", Account: "123456", Digit: "X",
})  // true
```

---

### IBAN

Utilitários para o IBAN (International Bank Account Number) brasileiro.

```go
import "github.com/brazilian-utils/go/iban"

// Validar IBAN
iban.IsValid("BR1500000000000010932840814P2")  // true

// Formatar IBAN
iban.Format("BR1500000000000010932840814P2")
// "BR15 0000 0000 0000 1093 2840 814P 2"

// Extrair informações do IBAN
info := iban.GetInfo("BR1500000000000010932840814P2")
fmt.Println(info.BankISPB, info.Account)  // "00000000" "0932840814"
```

---

### Credit Card

Validação de número de cartão de pagamento (crédito ou débito) pelo algoritmo de Luhn, sem detecção de bandeira.

```go
import "github.com/brazilian-utils/go/creditcard"

// Validar número de cartão
creditcard.IsValid("4111111111111111")     // true
creditcard.IsValid("4111 1111 1111 1111")  // true (com espaços)
creditcard.IsValid("4111111111111112")     // false (dígito de Luhn inválido)
```

---

### Pix Key

Identifica e normaliza uma chave Pix (CPF, CNPJ, email, telefone celular ou chave aleatória - EVP).

```go
import "github.com/brazilian-utils/go/pixkey"

// Validar chave Pix
pixkey.IsValid("123.456.789-09")  // true (CPF)
pixkey.IsValid("+5561912345678")  // true (telefone)

// Identificar e normalizar a chave
info := pixkey.GetInfo("123.456.789-09")
fmt.Println(info.Type, info.Value)  // "cpf" "12345678909"
```

---

### Pix Payload

Geração, leitura e validação do payload do BR Code Pix ("Pix copia e cola").

```go
import "github.com/brazilian-utils/go/pixpayload"

// Gerar um payload Pix estático
payload := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{
    Key:          "123e4567-e12b-12d1-a456-426655440000",
    MerchantName: "Fulano de Tal",
    MerchantCity: "Brasilia",
})

// Validar o payload gerado
pixpayload.IsValid(payload)  // true

// Extrair informações do payload
info := pixpayload.GetInfo(payload)
fmt.Println(info.Key, info.MerchantName)
// "123e4567-e12b-12d1-a456-426655440000" "Fulano de Tal"
```

---

### Area Code

Consulta aos 67 DDDs (códigos de área) brasileiros em uso, com o estado e a região correspondentes (Plano Geral de Numeração da Anatel).

```go
import "github.com/brazilian-utils/go/areacode"

// Consultar informações do DDD
info := areacode.GetInfo("11")
fmt.Println(info.StateCode, info.RegionCode)  // "SP" "SE"

// Listar os DDDs de um estado
areacode.ListByState("SP")  // [11 12 13 14 15 16 17 18 19]
```

---

### Municipality

Consulta aos municípios brasileiros publicados pelo IBGE.

```go
import "github.com/brazilian-utils/go/municipality"

// Buscar município pelo código do IBGE
m := municipality.GetByCode("3550308")
fmt.Println(m.Name, m.StateCode)  // "São Paulo" "SP"

// Listar municípios de um estado
df := municipality.List("DF")
fmt.Println(df[0].Name)  // "Brasília"
```

---

### State

Consulta às 27 unidades federativas brasileiras (26 estados + Distrito Federal).

```go
import "github.com/brazilian-utils/go/state"

// Buscar estado pelo código do IBGE
sp := state.GetByIbgeCode("35")
fmt.Println(sp.Code, sp.Name)  // "SP" "São Paulo"

// Nome a partir da sigla e sigla a partir do nome
state.GetNameByCode("RJ")          // "Rio de Janeiro"
state.GetCodeByName("São Paulo")   // "SP"

// Fuso horário do estado
state.GetTimezone("SP")  // "America/Sao_Paulo"
```

---

### VIN

Validação estrutural do VIN (Vehicle Identification Number / chassi) de um veículo: 17 caracteres, sem I/O/Q, e dígito verificador.

```go
import "github.com/brazilian-utils/go/vin"

// Validar VIN
vin.IsValid("1HGCM82633A004352")  // true
vin.IsValid("1HGCM8263IA004352")  // false (contém "I", letra não permitida)
```

---

### Text

Utilitários de texto com reconhecimento do português brasileiro: capitalização de nomes/endereços e remoção de acentos.

```go
import "github.com/brazilian-utils/go/text"

// Capitalizar nome/endereço
text.Capitalize("esponja de    aço 60G")  // "Esponja de Aço 60g"
text.Capitalize("fulano de tal")          // "Fulano de Tal"

// Remover acentos
text.RemoveAccents("São Paulo")  // "Sao Paulo"
```

---

### Number

Converte um número inteiro para sua forma por extenso em português.

```go
import "github.com/brazilian-utils/go/number"

// Converter número para texto
number.ConvertToWords(1235)  // "mil duzentos e trinta e cinco"
number.ConvertToWords(-3)    // "menos três"

// Forma feminina
number.ConvertToWords(2, number.ConvertToWordsOptions{Feminine: true})  // "duas"
```

---

### 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para enviar um Pull Request.

### 📄 Licença

Este projeto está licenciado sob a Licença MIT.

---

## English

### 📖 About

**Brazilian Utils for Go** is a comprehensive library that provides utilities for working with Brazilian-specific data formats and documents. It includes validators, formatters, and generators for various Brazilian documents and data types.

### 🚀 Installation

```shell
go get -u github.com/brazilian-utils/go
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
- **[Voter ID](#voter-id)** - Voter Registration ID

**Personal & professional documents:**
- **[CNS](#cns)** - National Health Card
- **[Certidão](#certidão)** - Civil Registry Certificate
- **[Passport](#passport)** - Brazilian Passport
- **[Professional Registration](#professional-registration)** - Professional Council Registration
- **[CBO](#cbo)** - Brazilian Occupation Classification

**Works/employer registries:**
- **[CEI](#cei)** - INSS-Specific Registry
- **[CNO](#cno)** - National Works Registry

**Tax registrations & fiscal codes:**
- **[CAEPF](#caepf)** - Individual Economic Activity Registry
- **[IE](#ie)** - State Tax Registration
- **[CFOP](#cfop)** - Tax Code for Operations and Services
- **[CNAE](#cnae)** - National Classification of Economic Activities
- **[CSOSN](#csosn)** - Simples Nacional Tax Situation Code
- **[CST](#cst)** - Tax Situation Code
- **[NCM](#ncm)** - Mercosul Common Nomenclature
- **[NFe Key](#nfe-key)** - Electronic Tax Document Access Key

**Financial & banking:**
- **[Bank](#bank)** - Brazilian Banks (COMPE)
- **[Bank Account](#bank-account)** - Brazilian Bank Account
- **[IBAN](#iban)** - Brazilian IBAN
- **[Credit Card](#credit-card)** - Payment Card Validation
- **[Pix Key](#pix-key)** - Pix Key
- **[Pix Payload](#pix-payload)** - Pix BR Code Payload

**Location:**
- **[Area Code](#area-code)** - Brazilian DDDs
- **[Municipality](#municipality)** - Brazilian Municipalities (IBGE)
- **[State](#state)** - Brazilian Federative Units

**Vehicles:**
- **[VIN](#vin)** - Vehicle Identification Number

**Text & numbers:**
- **[Text](#text)** - Portuguese Text Utilities
- **[Number](#number)** - Number to Words

---

### CPF

CPF (Cadastro de Pessoas Físicas) is the Brazilian individual taxpayer identification number.

```go
import "github.com/brazilian-utils/go/cpf"

// Validate CPF
cpf.IsValid("40364478829")  // true
cpf.IsValid("403.644.788-29")  // true
cpf.IsValid("00000000000")  // false (blacklisted)

// Format CPF
cpf.Format("40364478829")  // "403.644.788-29"
cpf.Format("403644788")  // "403.644.788" (incomplete)

// Generate random valid CPF
cpf.Generate()  // "12345678909" (random)

// Remove formatting symbols
cpf.RemoveSymbols("831.595.621-31")  // "83159562131"

// Keep only digits (normalize, capped to 11)
cpf.Parse("943.895.751-04")  // "94389575104"
```

---

### CNPJ

CNPJ (Cadastro Nacional da Pessoa Jurídica) is the Brazilian company identification number.

```go
import "github.com/brazilian-utils/go/cnpj"

// Validate CNPJ
cnpj.IsValid("11222333000181")  // true
cnpj.IsValid("11.222.333/0001-81")  // true
cnpj.IsValid("00000000000000")  // false (blacklisted)

// Format CNPJ
cnpj.Format("11222333000181")  // "11.222.333/0001-81"

// Generate random valid CNPJ
cnpj.Generate()  // "12345678000190" (random)

// Generate and validate an alphanumeric CNPJ (v2)
alnum := cnpj.Generate(cnpj.GenerateOptions{Version: 2})
cnpj.IsValid(alnum, cnpj.IsValidOptions{Version: 2})  // true

// Remove formatting symbols
cnpj.RemoveSymbols("10.799.163/9892-71")  // "10799163989271"

// Keep only digits (normalize, capped to 14)
cnpj.Parse("46.843.485/0001-86")  // "46843485000186"
```

---

### CEP

CEP (Código de Endereçamento Postal) is the Brazilian postal code.

```go
import "github.com/brazilian-utils/go/cep"

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

// Remove formatting symbols
cep.RemoveSymbols("91906-292")  // "91906292"

// Keep only digits (normalize, capped to 8)
cep.Parse("01001-000")  // "01001000"
```

---

### Phone

Utilities for Brazilian phone numbers (mobile and landline).

```go
import "github.com/brazilian-utils/go/phone"

// Validate phone numbers (mobile and landline accepted by default)
phone.IsValid("11994029275")  // true
phone.IsValid("1635014415")   // true

// Restrict the accepted kind
phone.IsValid("11994029275", phone.IsValidOptions{Kinds: []string{"mobile"}})    // true
phone.IsValid("11994029275", phone.IsValidOptions{Kinds: []string{"landline"}})  // false

// Validate a specific kind directly
phone.IsValidMobile("(11) 98765-4321")     // true
phone.IsValidLandline("(11) 3000-0000")    // true
phone.IsValidService("0800 123 4567")      // true (toll-free/public utility number)

// Format phone numbers
phone.Format("988887777")                                      // "98888-7777" (subscriber number only)
phone.Format("1130000000", phone.FormatOptions{Mask: "ddd"})    // "(11)3000-0000"

// Remove symbols
phone.RemoveSymbols("(11) 99402-9275")  // "11994029275"

// Remove international code
phone.RemoveInternationalDialingCode("5511994029275")  // "11994029275"

// Keep only digits (normalize, country code stripped when present)
phone.Parse("+55 11 98888-7777")  // "11988887777"

// Generate random phone
phone.Generate("mobile")    // random mobile
phone.Generate("landline")  // random landline
phone.Generate("")          // random mobile or landline
```

---

### Currency

Format and convert Brazilian Real (R$) values.

```go
import "github.com/brazilian-utils/go/currency"

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
import "github.com/brazilian-utils/go/boleto"

// Validate boleto digitable line (47 digits)
boleto.IsValid("34191790010104351004791020150008291070026000")  // true/false

// Format the digitable line
boleto.Format("00190000090114971860168524522114675860000102656")
// "00190.00009 01149.718601 68524.522114 6 75860000102656"

// Keep only digits (normalize)
boleto.Parse("00190.00009 01149.718601 68524.522114 6 75860000102656")
// "00190000090114971860168524522114675860000102656"

// Generate a random valid boleto (cobrança bancária)
boleto.Generate()

// Extract the bank, amount and due date
info := boleto.GetInfo("00190000090114971860168524522114675860000102656")
fmt.Println(info.BankCode, info.AmountCents, info.DueDate)
// "001" 102656 2018-07-15
```

---

### Email

Email address validation.

```go
import "github.com/brazilian-utils/go/email"

// Validate email
email.IsValid("user@example.com")  // true
email.IsValid("invalid.email")     // false
email.IsValid(".user@example.com") // false (starts with dot)
```

---

### PIS

PIS (Programa de Integração Social) is a Brazilian social integration program number.

```go
import "github.com/brazilian-utils/go/pis"

// Validate PIS
pis.IsValid("12345678901")  // true/false

// Format PIS
pis.Format("12345678901")  // "123.45678.90-1"

// Generate random valid PIS
pis.Generate()  // "12345678901" (random)

// Remove formatting symbols
pis.RemoveSymbols("559.84775.36-3")  // "55984775363"

// Keep only digits (normalize, capped to 11)
pis.Parse("123.45678.90-1")  // "12345678901"
```

---

### CNH

CNH (Carteira Nacional de Habilitação) is the Brazilian national driver's license.

```go
import "github.com/brazilian-utils/go/cnh"

// Validate CNH (11 digits)
cnh.IsValid("12345678901")  // true/false
cnh.IsValid("00000000000")  // false (blacklisted)
```

---

### RENAVAM

RENAVAM is the Brazilian national vehicle registration number.

```go
import "github.com/brazilian-utils/go/renavam"

// Validate RENAVAM (11 digits)
renavam.IsValid("12345678901")  // true/false
renavam.IsValid("11111111111")  // false (all same digit)
```

---

### License Plate

Utilities for Brazilian vehicle license plates (old format and Mercosul).

```go
import "github.com/brazilian-utils/go/licenseplate"

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

// Remove formatting symbols
licenseplate.RemoveSymbols("WDG1S21")  // "WDG1S21"

// Parse and normalize (upper-cased, capped to 7 characters)
licenseplate.Parse("abc-1234")  // "ABC1234"

// Generate random license plate
licenseplate.Generate("LLLNNNN")  // "ABC1234" (old format)
licenseplate.Generate("LLLNLNN")  // "ABC1D34" (Mercosul)
```

---

### Date

Convert Brazilian dates to Portuguese text.

```go
import "github.com/brazilian-utils/go/date"

// Convert date to text
date.ConvertDateToText("25/12/2024")  // "Vinte e cinco de Dezembro de dois mil e vinte e quatro"
date.ConvertDateToText("01/01/2000")  // "Primeiro de Janeiro de dois mil"

// Check whether a date is a holiday (national and, optionally, a state's)
date.IsHoliday(date.IsHolidayOptions{Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)})  // true

// List the national holidays of a year
date.GetHolidays(2024)

// Check whether it's a business day, and add/subtract business days
date.IsBusinessDay(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC))  // true (Tuesday)
date.AddBusinessDays(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 1)
// resolves to 2024-01-02 (skips the New Year holiday)

// Difference in business days between two dates
date.DifferenceInBusinessDays(
    time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
)  // 4
```

---

### Legal Nature

Legal Nature (Natureza Jurídica) codes from the Brazilian Federal Revenue.

```go
import "github.com/brazilian-utils/go/legalnature"

// Validate legal nature code
legalnature.IsValid("2062")  // true
legalnature.IsValid("206-2")  // true

// Get description
legalnature.GetDescription("2062")  // "Sociedade Empresária Limitada"
legalnature.GetDescription("2046")  // "Sociedade Anônima Aberta"

// Look up the full entry (code, description and category)
ln := legalnature.Get("2062")
fmt.Println(ln.Description, ln.Category.Code)  // "Sociedade Empresária Limitada" "2"

// List every code in force (92) or by CONCLA category
legalnature.List()
legalnature.ListByCategory("5")

// Format, parse and generate a random valid code
legalnature.Format("2062")  // "206-2"
legalnature.Parse("206-2")  // "2062"
legalnature.Generate()      // random code among those in force
```

---

### Legal Process

Brazilian legal process number utilities.

```go
import "github.com/brazilian-utils/go/legalprocess"

// Validate legal process (20 digits)
legalprocess.IsValid("12345670820231234567")  // true/false

// Format legal process
legalprocess.Format("12345670820231234567")  // "1234567-08.2023.1.23.4567"

// Generate random legal process
legalprocess.Generate(legalprocess.GenerateOptions{Court: 1})  // random, current year, segment 1

// Remove formatting symbols
legalprocess.RemoveSymbols("1666923-98.2026.9.26.4931")  // "16669239820269264931"

// Keep only digits (normalize, capped to 20)
legalprocess.Parse("0002080-25.2012.5.15.0049")  // "00020802520125150049"
```

---

### Voter ID

Voter ID (Título de Eleitor) is the Brazilian voter registration identification number.

```go
import "github.com/brazilian-utils/go/voterid"

// Validate Voter ID
voterid.IsValid("217633460930")  // true
voterid.IsValid("6908 4709 28 28")  // true (with spaces)
voterid.IsValid("123456789011")  // false (invalid check digits)

// Format Voter ID
voterid.Format("217633460930")  // "2176 3346 09 30"
voterid.Format("3244567800167")  // "3244 5678 00 16" (13 digits for SP/MG)

// Generate random valid Voter ID
voterid.Generate()       // "217633460930" (random, default ZZ - foreigner)
voterid.Generate("SP")   // "149426030183" (São Paulo)
voterid.Generate("MG")   // "033568860230" (Minas Gerais)
voterid.Generate("RJ")   // "858224120973" (Rio de Janeiro)
```

**Notes:**
- Voter IDs are typically 12 digits long
- Special cases: SP (01) and MG (02) can have 13 digits
- Valid UF codes: 01-28
- Two check digits calculated using modulo 11

---

### CNS

CNS (Cartão Nacional de Saúde) is the Brazilian national health card (SUS) number.

```go
import "github.com/brazilian-utils/go/cns"

// Validate CNS
cns.IsValid("123456789010000")     // true
cns.IsValid("123 4567 8901 0000")  // true (with spaces)
cns.IsValid("700000000000005")     // true (provisional card)

// Format CNS
cns.Format("123456789010001")  // "123 4567 8901 0001"

// Remove formatting
cns.Parse("898 0000 0004 3208")  // "898000000043208"
```

---

### Certidão

Validation, formatting and parsing of the matrícula (registration number) of a Brazilian civil registry certificate (birth, marriage, death, etc.), per art. 473 of the Código Nacional de Normas da Corregedoria Nacional de Justiça.

```go
import "github.com/brazilian-utils/go/certidao"

// Validate the matrícula (32 digits)
certidao.IsValid("10453901552013100012021000012321")  // true

// Format the matrícula
certidao.Format("10453901552013100012021000012321")
// "104539 01 55 2013 1 00012 021 0000123 21"

// Extract the matrícula's fields
info := certidao.GetInfo("104539 01 55 2013 1 00012 021 0000123 21")
fmt.Println(info.Year, info.Type)  // 2013 "birth"
```

---

### Passport

Utilities for the Brazilian passport number (2 letters followed by 6 digits).

```go
import "github.com/brazilian-utils/go/passport"

// Validate a passport number
passport.IsValid("AA111111")   // true
passport.IsValid("AB-123456")  // true (with symbols)

// Format/normalize
passport.Format("acd12736")  // "ACD12736"

// Generate a random valid passport number
passport.Generate()  // "AB123456" (random)
```

---

### Professional Registration

Structural check (digit count and UF only) of a professional council registration number: OAB, CRM, CRO, CRP and CRC. None of them has a published check digit.

```go
import "github.com/brazilian-utils/go/registroprofissional"

// Validate an OAB registration
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "123456/SP", Council: "OAB",
})  // true

// Restrict by state
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "123456/SP", Council: "OAB", State: "RJ",
})  // false

// Validate a CRC registration (with the O/P registration-type digit)
registroprofissional.IsValid(registroprofissional.IsValidParams{
    Value: "SP-123456/O-3", Council: "CRC",
})  // true
```

---

### CBO

CBO (Classificação Brasileira de Ocupações) is the official Brazilian occupation code table (CBO 2002).

```go
import "github.com/brazilian-utils/go/cbo"

// Validate a CBO code
cbo.IsValid("212405")   // true
cbo.IsValid("2124-05")  // true (masked)

// Look up the official title
c := cbo.Get("212405")
fmt.Println(c.Description)  // "Analista de desenvolvimento de sistemas"

// Remove formatting
cbo.Parse("2124-05")  // "212405"
```

---

### CEI

CEI (Cadastro Específico do INSS) is the works/employer registration number, now superseded by CNO (but still in use).

```go
import "github.com/brazilian-utils/go/cei"

// Validate a CEI
cei.IsValid("11.583.00249/85")  // true
cei.IsValid("115830024985")     // true

// Format a CEI
cei.Format("277297118187")  // "27.729.71181/87"

// Remove formatting
cei.Parse("27.729.71181/87")  // "277297118187"
```

---

### CNO

CNO (Cadastro Nacional de Obras) is the works registration number that replaced CEI, keeping the same digit rules.

```go
import "github.com/brazilian-utils/go/cno"

// Validate a CNO
cno.IsValid("110840168062")     // true
cno.IsValid("11.084.01680/62")  // true

// Format a CNO
cno.Format("111130137368")  // "11.113.01373/68"
```

---

### CAEPF

CAEPF (Cadastro de Atividade Econômica da Pessoa Física) is the individual economic activity registration number.

```go
import "github.com/brazilian-utils/go/caepf"

// Validate a CAEPF
caepf.IsValid("293.118.610/001-84")  // true
caepf.IsValid("29311861000184")      // true

// Format a CAEPF
caepf.Format("29311861000184")  // "293.118.610/001-84"

// Remove formatting
caepf.Parse("293.118.610/001-84")  // "29311861000184"
```

---

### IE

IE (Inscrição Estadual) is a taxpayer's state tax registration number. Validation is structural (digit count per state) only, with no check digit (see the package doc).

```go
import "github.com/brazilian-utils/go/ie"

// Validate an IE for a given state
ie.IsValid(ie.IsValidParams{Value: "123456789012", State: "SP"})   // true
ie.IsValid(ie.IsValidParams{Value: "12345678", State: "RJ"})       // true
ie.IsValid(ie.IsValidParams{Value: "P123456789012", State: "SP"})  // true (rural producer)
```

---

### CFOP

CFOP (Código Fiscal de Operações e Prestações) is the official tax code for goods/services operations (Convênio SINIEF s/nº 1970, consolidated Anexo II).

```go
import "github.com/brazilian-utils/go/cfop"

// Validate a CFOP code
cfop.IsValid("5102")   // true
cfop.IsValid("5.102")  // true (masked)

// Look up the official description
c := cfop.Get("5102")
fmt.Println(c.Description)

// Remove formatting
cfop.Parse("5.102")  // "5102"
```

---

### CNAE

CNAE (Classificação Nacional de Atividades Econômicas) is the official economic activity subclass code (CNAE-Subclasses 2.3).

```go
import "github.com/brazilian-utils/go/cnae"

// Validate a CNAE code
cnae.IsValid("6201501")    // true
cnae.IsValid("6201-5/01")  // true (masked)

// Format a CNAE code
cnae.Format("6201501")  // "6201-5/01"

// Look up the official description
c := cnae.Get("6201501")
fmt.Println(c.Description)  // "DESENVOLVIMENTO DE PROGRAMAS DE COMPUTADOR SOB ENCOMENDA"
```

---

### CSOSN

CSOSN (Código de Situação da Operação no Simples Nacional) is the official code for the 10 Simples Nacional tax situations (Ajuste SINIEF 07/2005).

```go
import "github.com/brazilian-utils/go/csosn"

// Validate a CSOSN code
csosn.IsValid("101")  // true
csosn.IsValid(101)    // true (number)
csosn.IsValid("999")  // false (not in the table)
```

---

### CST

CST (Código de Situação Tributária) is the official ICMS, IPI and PIS/COFINS tax situation code.

```go
import "github.com/brazilian-utils/go/cst"

// Validate a CST (any table)
cst.IsValid("110")  // true (ICMS)
cst.IsValid("07")   // true (PIS/COFINS)

// Restrict to a specific table
cst.IsValid("07", cst.IsValidOptions{Tax: "ipi"})  // false
```

---

### NCM

NCM (Nomenclatura Comum do Mercosul) is the official Siscomex goods classification code, from the table currently in force.

```go
import "github.com/brazilian-utils/go/ncm"

// Validate an NCM code
ncm.IsValid("22030000")    // true
ncm.IsValid("2203.00.00")  // true (masked)

// Format an NCM code
ncm.Format("84713012")  // "8471.30.12"
```

---

### NFe Key

Access key (chave de acesso) of a Brazilian electronic tax document: NF-e, NFC-e, CT-e, MDF-e, CT-e OS, GTV-e, BP-e, NF3e and NFCom.

```go
import "github.com/brazilian-utils/go/nfekey"

key := "35170458716523000119550010000000121000123458"

// Validate the access key
nfekey.IsValid(key)  // true

// Format the access key
nfekey.Format(key)
// "3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458"

// Extract the key's fields
info := nfekey.GetInfo(key)
fmt.Println(info.StateCode, info.Year, info.Model)  // "SP" 2017 "55"
```

---

### Bank

Lookups over the Brazilian banks that carry a COMPE code, published by Banco Central in the STR participants list.

```go
import "github.com/brazilian-utils/go/bank"

// Look a bank up by its COMPE code
b := bank.GetByCode("341")
fmt.Println(b.Name)  // "ITAÚ UNIBANCO S.A."

// Look a bank up by its ISPB
bank.GetByIspb("00000000")  // Banco do Brasil (code 001)

// List every bank with a COMPE code
bank.List()
```

---

### Bank Account

Validation of a Brazilian bank account: bank code, agency, account and check digit. This is a partial implementation (see the package doc for what's covered).

```go
import "github.com/brazilian-utils/go/bankaccount"

// Validate a bank account
bankaccount.IsValid(bankaccount.IsValidParams{
    BankCode: "341", Agency: "1234", Account: "123456", Digit: "X",
})  // true
```

---

### IBAN

Utilities for the Brazilian IBAN (International Bank Account Number).

```go
import "github.com/brazilian-utils/go/iban"

// Validate an IBAN
iban.IsValid("BR1500000000000010932840814P2")  // true

// Format an IBAN
iban.Format("BR1500000000000010932840814P2")
// "BR15 0000 0000 0000 1093 2840 814P 2"

// Extract the IBAN's fields
info := iban.GetInfo("BR1500000000000010932840814P2")
fmt.Println(info.BankISPB, info.Account)  // "00000000" "0932840814"
```

---

### Credit Card

Validation of a generic payment card number (credit or debit) via the Luhn algorithm, with no brand detection.

```go
import "github.com/brazilian-utils/go/creditcard"

// Validate a card number
creditcard.IsValid("4111111111111111")     // true
creditcard.IsValid("4111 1111 1111 1111")  // true (with spaces)
creditcard.IsValid("4111111111111112")     // false (bad Luhn digit)
```

---

### Pix Key

Identifies and normalizes a Pix key (CPF, CNPJ, email, mobile phone or a random EVP key).

```go
import "github.com/brazilian-utils/go/pixkey"

// Validate a Pix key
pixkey.IsValid("123.456.789-09")  // true (CPF)
pixkey.IsValid("+5561912345678")  // true (phone)

// Identify and normalize the key
info := pixkey.GetInfo("123.456.789-09")
fmt.Println(info.Type, info.Value)  // "cpf" "12345678909"
```

---

### Pix Payload

Generate, parse and validate a Pix BR Code payload ("Pix copia e cola").

```go
import "github.com/brazilian-utils/go/pixpayload"

// Generate a static Pix payload
payload := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{
    Key:          "123e4567-e12b-12d1-a456-426655440000",
    MerchantName: "Fulano de Tal",
    MerchantCity: "Brasilia",
})

// Validate the generated payload
pixpayload.IsValid(payload)  // true

// Extract the payload's fields
info := pixpayload.GetInfo(payload)
fmt.Println(info.Key, info.MerchantName)
// "123e4567-e12b-12d1-a456-426655440000" "Fulano de Tal"
```

---

### Area Code

Lookups over the 67 Brazilian DDDs (area codes) in use, with the matching state and region (Anatel Plano Geral de Numeração).

```go
import "github.com/brazilian-utils/go/areacode"

// Look up a DDD's info
info := areacode.GetInfo("11")
fmt.Println(info.StateCode, info.RegionCode)  // "SP" "SE"

// List the DDDs of a state
areacode.ListByState("SP")  // [11 12 13 14 15 16 17 18 19]
```

---

### Municipality

Lookups over the Brazilian municipalities published by IBGE.

```go
import "github.com/brazilian-utils/go/municipality"

// Look a municipality up by its IBGE code
m := municipality.GetByCode("3550308")
fmt.Println(m.Name, m.StateCode)  // "São Paulo" "SP"

// List the municipalities of a state
df := municipality.List("DF")
fmt.Println(df[0].Name)  // "Brasília"
```

---

### State

Lookups over the 27 Brazilian federative units (26 states + Distrito Federal).

```go
import "github.com/brazilian-utils/go/state"

// Look a state up by its IBGE code
sp := state.GetByIbgeCode("35")
fmt.Println(sp.Code, sp.Name)  // "SP" "São Paulo"

// Name from the code, and code from the name
state.GetNameByCode("RJ")          // "Rio de Janeiro"
state.GetCodeByName("São Paulo")   // "SP"

// The state's time zone
state.GetTimezone("SP")  // "America/Sao_Paulo"
```

---

### VIN

Structural validation of a vehicle's VIN (Vehicle Identification Number / chassi): 17 characters, no I/O/Q, and a check digit.

```go
import "github.com/brazilian-utils/go/vin"

// Validate a VIN
vin.IsValid("1HGCM82633A004352")  // true
vin.IsValid("1HGCM8263IA004352")  // false (contains "I", a disallowed letter)
```

---

### Text

Brazilian-Portuguese-aware text utilities: capitalizing names/addresses and removing accents.

```go
import "github.com/brazilian-utils/go/text"

// Capitalize a name/address
text.Capitalize("esponja de    aço 60G")  // "Esponja de Aço 60g"
text.Capitalize("fulano de tal")          // "Fulano de Tal"

// Remove accents
text.RemoveAccents("São Paulo")  // "Sao Paulo"
```

---

### Number

Writes an integer out in Brazilian Portuguese words ("por extenso").

```go
import "github.com/brazilian-utils/go/number"

// Convert a number to words
number.ConvertToWords(1235)  // "mil duzentos e trinta e cinco"
number.ConvertToWords(-3)    // "menos três"

// Feminine gender
number.ConvertToWords(2, number.ConvertToWordsOptions{Feminine: true})  // "duas"
```

---

### 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### 📄 License

This project is licensed under the MIT License.