package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	ValueRaw string  `xml:"Value" json:"-"`
	Value    float64 `xml:"-" json:"value"`
}

func ParseXML(filePath string) ([]Valute, error) {
	xmlFile, openError := os.Open(filePath)
	if openError != nil {
		return nil, fmt.Errorf("failed to open XML file %q: %w", filePath, openError)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var catalog ValCurs
	if decodeError := decoder.Decode(&catalog); decodeError != nil {
		return nil, fmt.Errorf("failed to decode XML from %q: %w", filePath, decodeError)
	}

	for index := range catalog.Valutes {
		normalizedValue := strings.ReplaceAll(catalog.Valutes[index].ValueRaw, ",", ".")
		parsedFloat, parseError := strconv.ParseFloat(normalizedValue, 64)
		if parseError != nil {
			return nil, fmt.Errorf("failed to parse value %q for currency %s: %w", catalog.Valutes[index].ValueRaw, catalog.Valutes[index].CharCode, parseError)
		}
		catalog.Valutes[index].Value = parsedFloat
	}

	return catalog.Valutes, nil
}

func SortByValueDesc(currencies []Valute) {
	sort.Slice(currencies, func(firstIndex, secondIndex int) bool {
		return currencies[firstIndex].Value > currencies[secondIndex].Value
	})
}

func SaveJSON(filePath string, currencies []Valute) error {
	jsonBytes, marshalError := json.MarshalIndent(currencies, "", "  ")
	if marshalError != nil {
		return fmt.Errorf("failed to marshal currencies to JSON: %w", marshalError)
	}

	if writeError := os.WriteFile(filePath, jsonBytes, 0644); writeError != nil {
		return fmt.Errorf("failed to write JSON to %q: %w", filePath, writeError)
	}

	return nil
}
