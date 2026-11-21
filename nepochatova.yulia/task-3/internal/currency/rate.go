package currency

import (
	"encoding/xml"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Rate struct {
	NumCode  int     `xml:"NumCode"  json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value"    json:"value"`
}

func (r *Rate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	numCode, err := parseNumCode(raw.NumCode)
	if err != nil {
		return err
	}

	value, err := parseCommaFloat(raw.Value)
	if err != nil {
		return err
	}

	*r = Rate{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(raw.CharCode),
		Value:    value,
	}
	return nil
}

func parseNumCode(s string) (int, error) {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0, fmt.Errorf("NumCode is empty")
	}
	n, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("invalid NumCode %q: %w", s, err)
	}
	return n, nil
}

func parseCommaFloat(s string) (float64, error) {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return 0, fmt.Errorf("Value is empty")
	}
	normalized := strings.ReplaceAll(trimmed, ",", ".")
	v, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid Value %q: %w", s, err)
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, fmt.Errorf("Value %q must be a finite number", s)
	}
	return v, nil
}

func SortByValueDescending(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}
