package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var (
	errMissingValue   = errors.New("missing currency value")
	errNonFiniteValue = errors.New("non-finite currency value")
)

type Currency struct {
	NumCode  numericCode `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    *rate       `json:"value"     xml:"Value"`
}

type currencyDocument struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Entries []Currency `xml:"Valute"`
}

type numericCode int

func (code *numericCode) UnmarshalText(text []byte) error {
	raw := strings.TrimSpace(string(text))
	if raw == "" {
		*code = 0

		return nil
	}

	number, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("invalid NumCode %q: %w", raw, err)
	}

	*code = numericCode(number)

	return nil
}

type rate float64

func (r *rate) UnmarshalText(text []byte) error {
	raw := strings.TrimSpace(string(text))
	normalized := strings.ReplaceAll(raw, ",", ".")

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("invalid Value %q: %w", raw, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("invalid Value %q: %w", raw, errNonFiniteValue)
	}

	*r = rate(value)

	return nil
}

func Load(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var document currencyDocument

	document.Entries = make([]Currency, 0)
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	currencies := document.Entries
	for _, currency := range currencies {
		if currency.Value == nil {
			return nil, fmt.Errorf("currency %q: %w", currency.CharCode, errMissingValue)
		}
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return *currencies[i].Value > *currencies[j].Value
	})

	return currencies, nil
}
