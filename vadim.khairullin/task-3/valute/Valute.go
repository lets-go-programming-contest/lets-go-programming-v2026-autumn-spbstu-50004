package valute

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Valute struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"-"`
}

func (val *Valute) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type tempValute struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var temp tempValute

	if err := decoder.DecodeElement(&temp, &start); err != nil {
		return fmt.Errorf("failed to decode Valute element: %w", err)
	}

	val.NumCode = temp.NumCode
	val.CharCode = temp.CharCode

	valueStr := strings.ReplaceAll(temp.Value, ",", ".")

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Valute value %q: %w", temp.Value, err)
	}

	val.Value = value

	return nil
}
