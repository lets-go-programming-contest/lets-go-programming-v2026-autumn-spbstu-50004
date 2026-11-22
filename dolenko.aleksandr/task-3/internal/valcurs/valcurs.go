package valcurs

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var errNoCurrencies = errors.New("XML contains no currencies")

type CurrencyValue float64

func (v *CurrencyValue) UnmarshalText(data []byte) error {
	value := strings.TrimSpace(string(data))
	value = strings.ReplaceAll(value, ",", ".")

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("parse currency value: %w", err)
	}

	*v = CurrencyValue(number)

	return nil
}

type Valute struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value"     xml:"Value"`
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

	if len(currencies.Valutes) == 0 {
		return ValCurs{}, errNoCurrencies
	}

	return currencies, nil
}

func SortByValue(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}
