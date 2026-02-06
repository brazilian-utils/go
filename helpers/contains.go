package helpers

// Contains verifies if a slice of string contains the given string
func Contains(h []string, n string) bool {
	for _, s := range h {
		if s == n {
			return true
		}
	}

	return false
}

// ContainsInt verifies if a slice of int contains the given int
func ContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}

	return false
}
