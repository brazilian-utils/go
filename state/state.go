// Package state provides lookups over the 27 Brazilian federative units
// (26 states + Distrito Federal).
package state

import (
	_ "embed"
	"encoding/json"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/brazilian-utils/go/text"
)

//go:embed data.json
var rawData []byte

// State is one Brazilian federative unit.
type State struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	RegionCode string `json:"regionCode"`
	RegionName string `json:"regionName"`
	IBGECode   int    `json:"ibgeCode"`
	Timezone   string `json:"timezone"`
}

type sourceFile struct {
	Data []State `json:"data"`
}

var (
	all         []State
	byCode      map[string]State
	byIBGECode  map[int]State
	byNameKey   map[string]State
)

var spaceRegex = regexp.MustCompile(`\s+`)

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = spaceRegex.ReplaceAllString(name, " ")
	name = text.RemoveAccents(name)
	return strings.ToLower(name)
}

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("state: failed to parse embedded data.json: " + err.Error())
	}
	all = src.Data

	byCode = make(map[string]State, len(all))
	byIBGECode = make(map[int]State, len(all))
	byNameKey = make(map[string]State, len(all))
	for _, s := range all {
		byCode[s.Code] = s
		byIBGECode[s.IBGECode] = s
		byNameKey[normalizeName(s.Name)] = s
	}

	sort.Slice(all, func(i, j int) bool { return all[i].Name < all[j].Name })
}

// List returns the 27 Brazilian federative units, sorted by name.
func List() []State {
	result := make([]State, len(all))
	copy(result, all)
	return result
}

// GetByIbgeCode returns the state whose 2-digit IBGE code matches code
// (a string or an int); returns nil when no state matches.
func GetByIbgeCode(code interface{}) *State {
	n, ok := toInt(code)
	if !ok {
		return nil
	}
	s, ok := byIBGECode[n]
	if !ok {
		return nil
	}
	copied := s
	return &copied
}

// GetCodeByName returns the two-letter code (sigla) of a state from its
// full name; the match ignores accents, case and surrounding/internal
// whitespace. Returns nil when no state matches.
func GetCodeByName(name string) *string {
	s, ok := byNameKey[normalizeName(name)]
	if !ok {
		return nil
	}
	code := s.Code
	return &code
}

// GetNameByCode returns the full name of a state from its two-letter
// code; the match ignores case and surrounding whitespace. Returns nil
// when no state matches.
func GetNameByCode(code string) *string {
	code = strings.ToUpper(strings.TrimSpace(code))
	s, ok := byCode[code]
	if !ok {
		return nil
	}
	name := s.Name
	return &name
}

// GetTimezone returns the IANA time zone of a state (the zone of its
// capital); the match ignores case and surrounding whitespace. Returns
// nil when no state matches.
func GetTimezone(stateCode string) *string {
	stateCode = strings.ToUpper(strings.TrimSpace(stateCode))
	s, ok := byCode[stateCode]
	if !ok {
		return nil
	}
	tz := s.Timezone
	return &tz
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
