package legalnature

import (
	_ "embed"
	"encoding/json"
	"sort"
)

//go:embed data.json
var rawData []byte

// Category is the CONCLA category a legal nature code belongs to (given by
// the first digit of the code).
type Category struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// LegalNature is one entry of the CONCLA/IBGE Natureza Jurídica 2021 table.
type LegalNature struct {
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
	Legacy      bool     `json:"legacy"`
	CurrentCode *string  `json:"currentCode,omitempty"`
}

type sourceFile struct {
	Data    []LegalNature `json:"data"`
	Retired []LegalNature `json:"retired"`
}

var (
	inForce      []LegalNature
	retired      []LegalNature
	byCode       map[string]LegalNature
	inForceByCode map[string]LegalNature
)

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("legalnature: failed to parse embedded data.json: " + err.Error())
	}

	inForce = src.Data
	retired = src.Retired

	byCode = make(map[string]LegalNature, len(inForce)+len(retired))
	inForceByCode = make(map[string]LegalNature, len(inForce))
	for _, ln := range inForce {
		byCode[ln.Code] = ln
		inForceByCode[ln.Code] = ln
	}
	for _, ln := range retired {
		byCode[ln.Code] = ln
	}

	sort.Slice(inForce, func(i, j int) bool { return inForce[i].Code < inForce[j].Code })
}
