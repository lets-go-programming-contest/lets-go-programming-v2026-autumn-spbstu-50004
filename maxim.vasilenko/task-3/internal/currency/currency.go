package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const dirPermissions = 0o755

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("failed to decode xml element: %w", err)
	}

	cleaned := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")

	val, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float %q: %w", raw, err)
	}

	*cf = CommaFloat(val)

	return nil
}

func (cf CommaFloat) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(float64(cf))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal float to json: %w", err)
	}

	return data, nil
}

type valCurs struct {
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

func ParseXML(filePath string) ([]Currency, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsedCurs valCurs
	if err := decoder.Decode(&parsedCurs); err != nil {
		return nil, fmt.Errorf("failed to decode xml: %w", err)
	}

	for idx := range parsedCurs.Valutes {
		parsedCurs.Valutes[idx].CharCode = strings.TrimSpace(parsedCurs.Valutes[idx].CharCode)
	}

	return parsedCurs.Valutes, nil
}

func SortByValueDesc(currencies []Currency) {
	sort.Slice(currencies, func(firstIdx, secondIdx int) bool {
		return currencies[firstIdx].Value > currencies[secondIdx].Value
	})
}

func SaveJSON(filePath string, currencies []Currency) error {
	dirPath := filepath.Dir(filePath)
	if err := os.MkdirAll(dirPath, dirPermissions); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer func() {
		_ = outFile.Close()
	}()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	return nil
}
