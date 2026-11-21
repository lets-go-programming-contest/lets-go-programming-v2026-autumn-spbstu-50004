package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func ParseXML(data []byte) (ValCurs, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return ValCurs{}, fmt.Errorf("failed to decode XML: %w", err)
	}

	return valCurs, nil
}

func parseValue(value string) (float64, error) {
	normalized := strings.ReplaceAll(value, ",", ".")

	result, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse value: %w", err)
	}

	return result, nil
}

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ToResults(valutes []Valute) ([]CurrencyResult, error) {
	results := make([]CurrencyResult, 0, len(valutes))
	for _, valute := range valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			return nil, fmt.Errorf("invalid NumCode %q: %w", valute.NumCode, err)
		}

		value, err := parseValue(valute.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid Value %q: %w", valute.Value, err)
		}

		results = append(results, CurrencyResult{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return results, nil
}

func SortByValueDescending(results []CurrencyResult) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})
}
