package rates

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlValCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Items   []xmlValute `xml:"Valute"`
}

type xmlValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func FetchSortedRates(xmlPath string) ([]Rate, error) {
	file, err := os.Open(filepath.Clean(xmlPath))
	if err != nil {
		return nil, fmt.Errorf("open xml source: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs xmlValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml structure: %w", err)
	}

	rates := make([]Rate, 0, len(curs.Items))
	for _, item := range curs.Items {
		rate, err := convertItem(item)
		if err != nil {
			return nil, err
		}

		rates = append(rates, rate)
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}

func convertItem(item xmlValute) (Rate, error) {
	numStr := strings.TrimSpace(item.NumCode)
	numCode := 0

	if numStr != "" {
		parsedNum, err := strconv.Atoi(numStr)
		if err != nil {
			return Rate{}, fmt.Errorf("invalid num_code %q: %w", numStr, err)
		}

		numCode = parsedNum
	}

	valStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return Rate{}, fmt.Errorf("invalid currency value %q: %w", item.Value, err)
	}

	return Rate{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(item.CharCode),
		Value:    val,
	}, nil
}
