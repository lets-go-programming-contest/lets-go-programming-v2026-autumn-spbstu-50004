package cbr

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

type CurrencyCode int

func (c *CurrencyCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode currency code: %w", err)
	}

	val, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil {
		return fmt.Errorf("parse currency code %q: %w", raw, err)
	}

	*c = CurrencyCode(val)

	return nil
}

type CurrencyValue float64

func (c *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode currency value: %w", err)
	}

	normalized := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")

	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse currency value %q: %w", raw, err)
	}

	*c = CurrencyValue(val)

	return nil
}

type Currency struct {
	NumCode  CurrencyCode  `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value"     xml:"Value"`
}

type valCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

func ParseXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("read cbr xml: %w", err)
	}

	var curs valCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode cbr xml: %w", err)
	}

	rates := curs.Valutes

	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}

func WriteJSON(path string, rates []Currency) error {
	data, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json data: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
