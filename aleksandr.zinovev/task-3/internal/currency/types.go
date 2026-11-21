package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	errEmptyValue     = errors.New("currency value is empty")
	errNonFiniteValue = errors.New("currency value must be finite")
)

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode value element: %w", err)
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return errEmptyValue
	}

	normalized := strings.ReplaceAll(trimmed, ",", ".")

	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", raw, err)
	}

	if math.IsNaN(val) || math.IsInf(val, 0) {
		return errNonFiniteValue
	}

	*cf = CommaFloat(val)

	return nil
}

type OptionalInt int

func (oi *OptionalInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode num code element: %w", err)
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		*oi = 0

		return nil
	}

	val, err := strconv.Atoi(trimmed)
	if err != nil {
		return fmt.Errorf("parse int %q: %w", raw, err)
	}

	*oi = OptionalInt(val)

	return nil
}

type Valute struct {
	NumCode  OptionalInt `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    CommaFloat  `json:"value"     xml:"Value"`
}
