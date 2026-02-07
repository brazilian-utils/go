package voterid

import (
	"fmt"
	"math/rand"
	"strings"
)

// State to federative union code mapping
var stateToFUCode = map[string]string{
	"SP": "01",
	"MG": "02",
	"RJ": "03",
	"RS": "04",
	"BA": "05",
	"PR": "06",
	"CE": "07",
	"PE": "08",
	"SC": "09",
	"GO": "10",
	"MA": "11",
	"PB": "12",
	"PA": "13",
	"ES": "14",
	"PI": "15",
	"RN": "16",
	"AL": "17",
	"MT": "18",
	"MS": "19",
	"DF": "20",
	"SE": "21",
	"AM": "22",
	"RO": "23",
	"AC": "24",
	"AP": "25",
	"RR": "26",
	"TO": "27",
	"ZZ": "28", // For voter IDs issued to foreigners
}

// Generate generates a random valid Brazilian voter ID
// The federativeUnion parameter accepts state abbreviations (e.g., "SP", "MG", "RJ")
// Default is "ZZ" for voter IDs issued to foreigners
// Returns empty string if the state code is invalid
func Generate(federativeUnion ...string) string {
	state := "ZZ"
	if len(federativeUnion) > 0 {
		state = strings.ToUpper(federativeUnion[0])
	}

	// Get the federative union code
	ufCode, exists := stateToFUCode[state]
	if !exists {
		return ""
	}

	// Validate the federative union code
	if !isFederativeUnionValid(ufCode) {
		return ""
	}

	// Generate random 8-digit sequential number
	sequentialNumber := fmt.Sprintf("%08d", rand.Intn(99999999))

	// Calculate verifying digits
	vd1 := calculateVD1(sequentialNumber, ufCode)
	vd2 := calculateVD2(ufCode, vd1)

	return fmt.Sprintf("%s%s%d%d", sequentialNumber, ufCode, vd1, vd2)
}
