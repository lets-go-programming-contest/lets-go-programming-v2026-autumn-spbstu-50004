package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	ValueRaw string  `json:"-"         xml:"Value"`
	Value    float64 `json:"value"     xml:"-"`
}

func ParseXML(filePath string) ([]Valute, error) {
	xmlFile, openError := os.Open(filePath)
	if openError != nil {
		return nil, fmt.Errorf("failed to open XML file %q: %w", filePath, openError)
	}

	defer func() {
		if closeError := xmlFile.Close(); closeError != nil {
			log.Println(fmt.Errorf("failed to close XML file %q: %w", filePath, closeError))
		}
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var catalog ValCurs

	if decodeError := decoder.Decode(&catalog); decodeError != nil {
		return nil, fmt.Errorf("failed to decode XML from %q: %w", filePath, decodeError)
	}

	for index := range catalog.Valutes {
		valute := &catalog.Valutes[index]
		normalizedValue := strings.ReplaceAll(valute.ValueRaw, ",", ".")

		parsedFloat, parseError := strconv.ParseFloat(normalizedValue, 64)
		if parseError != nil {
			return nil, fmt.Errorf("failed to parse value %q for %s: %w", valute.ValueRaw, valute.CharCode, parseError)
		}

		valute.Value = parsedFloat
	}

	return catalog.Valutes, nil
}

func SortByValueDesc(currencies []Valute) {
	sort.Slice(currencies, func(firstIndex, secondIndex int) bool {
		return currencies[firstIndex].Value > currencies[secondIndex].Value
	})
}

func SaveJSON(filePath string, currencies []Valute) error {
	outputDir := filepath.Dir(filePath)

	if mkdirError := os.MkdirAll(outputDir, 0o750); mkdirError != nil {
		return fmt.Errorf("failed to create output directory %q: %w", outputDir, mkdirError)
	}

	jsonBytes, marshalError := json.MarshalIndent(currencies, "", "  ")
	if marshalError != nil {
		return fmt.Errorf("failed to marshal currencies to JSON: %w", marshalError)
	}

	if writeError := os.WriteFile(filePath, jsonBytes, 0o600); writeError != nil {
		return fmt.Errorf("failed to write JSON to %q: %w", filePath, writeError)
	}

	return nil
}
