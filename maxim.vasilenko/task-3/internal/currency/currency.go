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

type valCurs struct {
	Valutes []rawValute `xml:"Valute"`
}

type rawValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

	currencies := make([]Currency, 0, len(parsedCurs.Valutes))

	for _, item := range parsedCurs.Valutes {
		valStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		parsedVal, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse currency value %q: %w", item.Value, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    parsedVal,
		})
	}

	return currencies, nil
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
