// Package areacode provides lookups over the 67 Brazilian DDDs (area
// codes) in use under the Anatel Plano Geral de Numeração.
package areacode

import (
	_ "embed"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.json
var rawData []byte

// Info is the state/region information for a Brazilian DDD (area code).
type Info struct {
	AreaCode   int      `json:"areaCode"`
	StateCode  string   `json:"stateCode"`
	StateName  string   `json:"stateName"`
	RegionCode string   `json:"regionCode"`
	RegionName string   `json:"regionName"`
	StateCodes []string `json:"stateCodes"`
}

type sourceFile struct {
	Data []Info `json:"data"`
}

var (
	all      []Info
	byCode   map[int]Info
	byState  map[string][]int
)

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("areacode: failed to parse embedded data.json: " + err.Error())
	}
	all = src.Data

	byCode = make(map[int]Info, len(all))
	byState = make(map[string][]int)
	for _, info := range all {
		byCode[info.AreaCode] = info
		for _, sc := range info.StateCodes {
			byState[sc] = append(byState[sc], info.AreaCode)
		}
	}
	for sc := range byState {
		sort.Ints(byState[sc])
	}
}

// GetInfo returns the state and region a DDD (area code) belongs to.
// areaCode may be a string or a non-negative integer. Returns nil for a
// DDD not in use, a negative or non-integer number, or anything else.
func GetInfo(areaCode interface{}) *Info {
	n, ok := toInt(areaCode)
	if !ok || n < 0 {
		return nil
	}
	info, ok := byCode[n]
	if !ok {
		return nil
	}
	copied := info
	return &copied
}

// ListByState returns every DDD (area code) that serves a state, sorted
// ascending. stateCode is matched case-insensitively. Returns an empty
// list when stateCode is not a Brazilian state.
func ListByState(stateCode string) []int {
	codes, ok := byState[strings.ToUpper(strings.TrimSpace(stateCode))]
	if !ok {
		return []int{}
	}
	result := make([]int, len(codes))
	copy(result, codes)
	return result
}

func toInt(value interface{}) (int, bool) {
	switch v := value.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case float64:
		return int(v), true
	case string:
		if v == "" {
			return 0, false
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, false
		}
		return n, true
	default:
		return 0, false
	}
}
