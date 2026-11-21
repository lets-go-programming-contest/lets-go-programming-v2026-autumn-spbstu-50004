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

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

type CurCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"Value"`
}

func (c *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("error decoding currency element: %w", err)
	}

	numCode := 0

	if s := strings.TrimSpace(raw.NumCode); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("wrong NumCode %q: %w", raw.NumCode, err)
		}

		numCode = n
	}

	value := 0.0

	if s := strings.ReplaceAll(strings.TrimSpace(raw.Value), ",", "."); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fmt.Errorf("wrong Value %q: %w", raw.Value, err)
		}

		value = v
	}

	c.NumCode = numCode
	c.CharCode = strings.TrimSpace(raw.CharCode)
	c.Value = value

	return nil
}

func Process(inputFile, outputFile string) error {
	curCurs, err := decodeXML(inputFile)
	if err != nil {
		return err
	}

	sort.Slice(curCurs.Currencies, func(i, j int) bool {
		return curCurs.Currencies[i].Value > curCurs.Currencies[j].Value
	})

	if err := saveToJSON(outputFile, curCurs.Currencies); err != nil {
		return err
	}

	return nil
}

func decodeXML(inputFile string) (*CurCurs, error) {
	xmlData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("error reading XML-file: %w", err)
	}

	var curCurs CurCurs

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&curCurs); err != nil {
		return nil, fmt.Errorf("error decrypting XML-file: %w", err)
	}

	return &curCurs, nil
}

func saveToJSON(outputFile string, currencies []Currency) error {
	dir := filepath.Dir(outputFile)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("error dir creation %q: %w", dir, err)
	}

	jsonData, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("error forming JSON: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
