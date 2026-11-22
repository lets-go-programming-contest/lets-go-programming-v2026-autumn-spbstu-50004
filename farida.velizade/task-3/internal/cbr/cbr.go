package cbr

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
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
