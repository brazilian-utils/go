package pixpayload

import "strconv"

// tlv is one Tag-Length-Value entry of an EMV QR Code payload.
type tlv struct {
	ID    string
	Value string
}

// parseTLV parses a flat sequence of {2-digit ID}{2-digit decimal
// length}{value} entries. Returns nil when the string is malformed.
func parseTLV(s string) []tlv {
	var result []tlv
	i := 0
	for i < len(s) {
		if i+4 > len(s) {
			return nil
		}
		id := s[i : i+2]
		lengthStr := s[i+2 : i+4]
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length < 0 {
			return nil
		}
		start := i + 4
		end := start + length
		if end > len(s) {
			return nil
		}
		result = append(result, tlv{ID: id, Value: s[start:end]})
		i = end
	}
	return result
}

func find(entries []tlv, id string) (string, bool) {
	for _, e := range entries {
		if e.ID == id {
			return e.Value, true
		}
	}
	return "", false
}

// crc16CCITT computes the CRC-16/CCITT-FALSE checksum (poly 0x1021,
// init 0xFFFF, no reflection, no final XOR) used by the EMV QR Code
// standard.
func crc16CCITT(s string) uint16 {
	var crc uint16 = 0xFFFF
	for i := 0; i < len(s); i++ {
		crc ^= uint16(s[i]) << 8
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc = crc << 1
			}
		}
	}
	return crc
}
