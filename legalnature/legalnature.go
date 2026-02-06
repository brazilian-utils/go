package legalnature

import "github.com/brazilian-utils/brutils-go/helpers"

// legalNatureTable maps official 4-digit codes to their descriptions,
// sourced from the Tabela de Natureza Jurídica (RFB).
var legalNatureTable = map[string]string{
	// 1. ADMINISTRAÇÃO PÚBLICA
	"1015": "Órgão Público do Poder Executivo Federal",
	"1023": "Órgão Público do Poder Executivo Estadual ou do Distrito Federal",
	"1031": "Órgão Público do Poder Executivo Municipal",
	"1040": "Órgão Público do Poder Legislativo Federal",
	"1058": "Órgão Público do Poder Legislativo Estadual ou do Distrito Federal",
	"1066": "Órgão Público do Poder Legislativo Municipal",
	"1074": "Órgão Público do Poder Judiciário Federal",
	"1082": "Órgão Público do Poder Judiciário Estadual",
	"1104": "Autarquia Federal",
	"1112": "Autarquia Estadual ou do Distrito Federal",
	"1120": "Autarquia Municipal",
	"1139": "Fundação Federal",
	"1147": "Fundação Estadual ou do Distrito Federal",
	"1155": "Fundação Municipal",
	"1163": "Órgão Público Autônomo da União",
	"1171": "Órgão Público Autônomo Estadual ou do Distrito Federal",
	"1180": "Órgão Público Autônomo Municipal",
	// 2. ENTIDADES EMPRESARIAIS
	"2011": "Empresa Pública",
	"2038": "Sociedade de Economia Mista",
	"2046": "Sociedade Anônima Aberta",
	"2054": "Sociedade Anônima Fechada",
	"2062": "Sociedade Empresária Limitada",
	"2076": "Sociedade Empresária em Nome Coletivo",
	"2089": "Sociedade Empresária em Comandita Simples",
	"2097": "Sociedade Empresária em Comandita por Ações",
	"2100": "Sociedade Mercantil de Capital e Indústria (extinta pelo NCC/2002)",
	"2127": "Sociedade Empresária em Conta de Participação",
	"2135": "Empresário (Individual)",
	"2143": "Cooperativa",
	"2151": "Consórcio de Sociedades",
	"2160": "Grupo de Sociedades",
	"2178": "Estabelecimento, no Brasil, de Sociedade Estrangeira",
	"2194": "Estabelecimento, no Brasil, de Empresa Binacional Argentino-Brasileira",
	"2208": "Entidade Binacional Itaipu",
	"2216": "Empresa Domiciliada no Exterior",
	"2224": "Clube/Fundo de Investimento",
	"2232": "Sociedade Simples Pura",
	"2240": "Sociedade Simples Limitada",
	"2259": "Sociedade em Nome Coletivo",
	"2267": "Sociedade em Comandita Simples",
	"2275": "Sociedade Simples em Conta de Participação",
	"2305": "Empresa Individual de Responsabilidade Limitada",
	// 3. ENTIDADES SEM FINS LUCRATIVOS
	"3034": "Serviço Notarial e Registral (Cartório)",
	"3042": "Organização Social",
	"3050": "Organização da Sociedade Civil de Interesse Público (Oscip)",
	"3069": "Outras Formas de Fundações Mantidas com Recursos Privados",
	"3077": "Serviço Social Autônomo",
	"3085": "Condomínio Edilícios",
	"3093": "Unidade Executora (Programa Dinheiro Direto na Escola)",
	"3107": "Comissão de Conciliação Prévia",
	"3115": "Entidade de Mediação e Arbitragem",
	"3123": "Partido Político",
	"3131": "Entidade Sindical",
	"3204": "Estabelecimento, no Brasil, de Fundação ou Associação Estrangeiras",
	"3212": "Fundação ou Associação Domiciliada no Exterior",
	"3999": "Outras Formas de Associação",
	// 4. PESSOAS FÍSICAS
	"4014": "Empresa Individual Imobiliária",
	"4022": "Segurado Especial",
	"4081": "Contribuinte individual",
	// 5. ORGANIZAÇÕES INTERNACIONAIS E OUTRAS INSTITUIÇÕES EXTRATERRITORIAIS
	"5002": "Organização Internacional e Outras Instituições Extraterritoriais",
}

// normalize strips non-digit characters and validates the code is 4 digits.
func normalize(code string) string {
	digits := helpers.OnlyNumbers(code)
	if len(digits) != 4 {
		return ""
	}
	return digits
}

// IsValid checks if a string corresponds to a valid Natureza Jurídica code.
// Accepts formats "NNNN" or "NNN-N".
func IsValid(code string) bool {
	normalized := normalize(code)
	if normalized == "" {
		return false
	}
	_, ok := legalNatureTable[normalized]
	return ok
}

// GetDescription retrieves the description of a Natureza Jurídica code.
// Returns empty string if the code is not found.
func GetDescription(code string) string {
	normalized := normalize(code)
	if normalized == "" {
		return ""
	}
	return legalNatureTable[normalized]
}

// ListAll returns a copy of the full Natureza Jurídica table.
func ListAll() map[string]string {
	result := make(map[string]string, len(legalNatureTable))
	for k, v := range legalNatureTable {
		result[k] = v
	}
	return result
}
