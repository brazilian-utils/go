// Package pixpayload generates, parses and validates a Pix BR Code
// payload (the string behind a Pix QR Code / "Pix copia e cola").
package pixpayload

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brazilian-utils/go/pixkey"
	"github.com/brazilian-utils/go/text"
)

const pixGUI = "br.gov.bcb.pix"

// Info is the information parsed from a Pix BR Code payload.
type Info struct {
	Key               string
	URL               string
	MerchantName      string
	MerchantCity      string
	Amount            *float64
	TxID              string
	Description       string
	PointOfInitiation string // "static" or "dynamic"
}

// findMerchantAccountInfo scans the root-level tags (26 to 51) for the
// one that carries the "br.gov.bcb.pix" GUI, and returns its sub-fields.
func findMerchantAccountInfo(root []tlv) []tlv {
	for _, entry := range root {
		id, err := strconv.Atoi(entry.ID)
		if err != nil || id < 26 || id > 51 {
			continue
		}
		sub := parseTLV(entry.Value)
		if sub == nil {
			continue
		}
		if gui, ok := find(sub, "00"); ok && gui == pixGUI {
			return sub
		}
	}
	return nil
}

// IsValid validates a Pix BR Code payload: the TLV structure, the
// CRC-16 and the mandatory objects. The key itself is not checked (use
// pixkey.IsValid).
func IsValid(value string) bool {
	return parse(value) != nil
}

// GetInfo parses a Pix BR Code payload into its fields; returns nil for
// anything IsValid rejects.
func GetInfo(value string) *Info {
	return parse(value)
}

func parse(value string) *Info {
	if len(value) < 4 {
		return nil
	}

	crcValue, ok := extractCRC(value)
	if !ok {
		return nil
	}
	withoutCRC := value[:len(value)-4]
	computed := crc16CCITT(withoutCRC)
	if !strings.EqualFold(fmt.Sprintf("%04X", computed), crcValue) {
		return nil
	}

	root := parseTLV(value)
	if root == nil {
		return nil
	}

	// Mandatory objects: format indicator, category code, currency,
	// country, merchant name and city.
	for _, id := range []string{"00", "52", "53", "58", "59", "60"} {
		if _, ok := find(root, id); !ok {
			return nil
		}
	}

	if poi, ok := find(root, "01"); ok && poi != "11" && poi != "12" {
		return nil
	}

	mai := findMerchantAccountInfo(root)
	if mai == nil {
		return nil
	}
	key, hasKey := find(mai, "01")
	url, hasURL := find(mai, "25")
	if hasKey == hasURL { // must have exactly one
		return nil
	}

	merchantName, _ := find(root, "59")
	merchantCity, _ := find(root, "60")

	info := &Info{
		MerchantName: merchantName,
		MerchantCity: merchantCity,
	}
	if hasKey {
		info.Key = key
		info.PointOfInitiation = "static"
	} else {
		info.URL = url
		info.PointOfInitiation = "dynamic"
	}
	if poi, ok := find(root, "01"); ok && poi == "12" {
		info.PointOfInitiation = "dynamic"
	}

	// With a PSP location, amount and txid are ignored.
	if info.PointOfInitiation != "dynamic" || !hasURL {
		if amountStr, ok := find(root, "54"); ok {
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				return nil
			}
			info.Amount = &amount
		}

		if addTemplate, ok := find(root, "62"); ok {
			sub := parseTLV(addTemplate)
			if txid, ok := find(sub, "05"); ok && txid != "***" {
				info.TxID = txid
			}
		}
	}

	return info
}

func extractCRC(value string) (string, bool) {
	if len(value) < 8 {
		return "", false
	}
	tag := value[len(value)-8 : len(value)-6]
	length := value[len(value)-6 : len(value)-4]
	if tag != "63" || length != "04" {
		return "", false
	}
	return strings.ToUpper(value[len(value)-4:]), true
}

// GeneratePixPayloadParams carries the fields of a payload to build.
// Exactly one of Key (static) or URL (dynamic) must be given.
type GeneratePixPayloadParams struct {
	Key          string
	URL          string
	MerchantName string
	MerchantCity string
	Amount       *float64
	TxID         string
	Description  string
}

// Generate builds a Pix BR Code payload; returns "" when neither or both
// of Key/URL are given, or the amount is invalid.
func Generate(params GeneratePixPayloadParams) string {
	hasKey := params.Key != ""
	hasURL := params.URL != ""
	if hasKey == hasURL {
		return ""
	}
	if hasKey && !pixkey.IsValid(params.Key) {
		return ""
	}
	if params.Amount != nil {
		rounded := float64(int64(*params.Amount*100+0.5)) / 100
		if rounded != *params.Amount || rounded == 0 {
			return ""
		}
	}

	merchantName := truncate(text.RemoveAccents(params.MerchantName), 25)
	merchantCity := truncate(text.RemoveAccents(params.MerchantCity), 15)

	var mai strings.Builder
	mai.WriteString(field("00", pixGUI))
	if hasKey {
		mai.WriteString(field("01", params.Key))
	} else {
		mai.WriteString(field("25", params.URL))
	}

	var b strings.Builder
	b.WriteString(field("00", "01"))
	if hasURL {
		b.WriteString(field("01", "12"))
	}
	b.WriteString(field("26", mai.String()))
	b.WriteString(field("52", "0000"))
	b.WriteString(field("53", "986"))
	if params.Amount != nil {
		b.WriteString(field("54", strconv.FormatFloat(*params.Amount, 'f', 2, 64)))
	}
	b.WriteString(field("58", "BR"))
	b.WriteString(field("59", merchantName))
	b.WriteString(field("60", merchantCity))

	if !hasURL {
		txid := params.TxID
		if txid == "" {
			txid = "***"
		}
		var add strings.Builder
		add.WriteString(field("05", txid))
		if params.Description != "" {
			add.WriteString(field("02", truncate(text.RemoveAccents(params.Description), 99)))
		}
		b.WriteString(field("62", add.String()))
	}

	withoutCRC := b.String() + "6304"
	crc := crc16CCITT(withoutCRC)
	return withoutCRC + fmt.Sprintf("%04X", crc)
}

func field(id, value string) string {
	return fmt.Sprintf("%s%02d%s", id, len(value), value)
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max]
	}
	return s
}
