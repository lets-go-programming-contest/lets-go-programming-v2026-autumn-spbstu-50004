package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	errEmptyValue     = errors.New("currency rate value cannot be empty")
	errNonFiniteValue = errors.New("currency rate value must be a finite number")
)

type Rate struct {
	NumCode  NumCode `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Value   `json:"value"     xml:"Value"`
}

type NumCode int

func (n *NumCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return fmt.Errorf("decode num_code element: %w", err)
	}

	raw = strings.TrimSpace(raw)
	if raw == "" {
		*n = 0

		return nil
	}

	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("invalid num_code %q: %w", raw, err)
	}

	*n = NumCode(parsed)

	return nil
}

type Value float64

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	err := d.DecodeElement(&raw, &start)
	if err != nil {
		return fmt.Errorf("decode value element: %w", err)
	}

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return errEmptyValue
	}

	parsed, err := strconv.ParseFloat(strings.ReplaceAll(raw, ",", "."), 64)
	if err != nil {
		return fmt.Errorf("invalid value %q: %w", raw, err)
	}

	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return errNonFiniteValue
	}

	*v = Value(parsed)

	return nil
}

func SortByValueDescending(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}
