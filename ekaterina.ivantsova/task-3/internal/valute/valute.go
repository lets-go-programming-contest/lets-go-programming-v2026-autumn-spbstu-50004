package valute

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
	errEmptyCurrencyValue   = errors.New("currency Value fields is empty")
	errNonFiniteCurrencyVal = errors.New("currency Value must be a finite number")
)

type RateValue float64

type Item struct {
	NumCode  int       `json:"num_code" xml:"NumCode"`
	CharCode string    `json:"char_code" xml:"CharCode"`
	Value    RateValue `json:"value" xml:"Value"`
}

func (value *RateValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var rawText string

	err := decoder.DecodeElement(&rawText, &start)
	if err != nil {
		return fmt.Errorf("decode Value element: %w", err)
	}

	cleanText := strings.ReplaceAll(strings.TrimSpace(rawText), ",", ".")
	if cleanText == "" {
		return errEmptyCurrencyValue
	}

	parsed, err := strconv.ParseFloat(cleanText, 64)
	if err != nil {
		return fmt.Errorf("invalid Value %q: %w", rawText, err)
	}

	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return errNonFiniteCurrencyVal
	}

	*value = RateValue(parsed)

	return nil
}

func SortByValueDesc(items []Item) {
	sort.SliceStable(items, func(left, right int) bool {
		return items[left].Value > items[right].Value
	})
}
