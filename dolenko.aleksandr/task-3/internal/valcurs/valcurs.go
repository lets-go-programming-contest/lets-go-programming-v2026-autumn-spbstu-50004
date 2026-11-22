package valcurs

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func Load(path string) (ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ValCurs{}, fmt.Errorf("read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var currencies ValCurs

	err = decoder.Decode(&currencies)
	if err != nil {
		return ValCurs{}, fmt.Errorf("decode XML: %w", err)
	}

	return currencies, nil
}
