package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/SoundwaveAI/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

var (
	errBlankXMLPath   = errors.New("xml file path cannot be blank")
	errInvalidRootTag = errors.New("unexpected root XML tag, expected ValCurs")
	errNoValutesFound = errors.New("XML document contains no Valute entries")
	errEmptyRateValue = errors.New("currency rate value cannot be empty")
	errNonFiniteValue = errors.New("currency rate value must be a finite number")
)

type cbrCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Entries []cbrValute `xml:"Valute"`
}

type cbrValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ExtractRates(xmlPath string) ([]currency.Rate, error) {
	trimmedPath := strings.TrimSpace(xmlPath)
	if trimmedPath == "" {
		return nil, errBlankXMLPath
	}

	cleanPath := filepath.Clean(trimmedPath)

	raw, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", cleanPath, err)
	}

	var curs cbrCurs

	decoder := xml.NewDecoder(bytes.NewReader(raw))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("decode XML from %q: %w", cleanPath, err)
	}

	if curs.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("%w: got %q", errInvalidRootTag, curs.XMLName.Local)
	}

	if len(curs.Entries) == 0 {
		return nil, errNoValutesFound
	}

	rates := make([]currency.Rate, 0, len(curs.Entries))

	for idx, entry := range curs.Entries {
		rate, convErr := convertEntry(entry, idx+1)
		if convErr != nil {
			return nil, convErr
		}

		rates = append(rates, rate)
	}

	return rates, nil
}

func convertEntry(entry cbrValute, row int) (currency.Rate, error) {
	numCode := 0
	rawCode := strings.TrimSpace(entry.NumCode)

	if rawCode != "" {
		parsedCode, err := strconv.Atoi(rawCode)
		if err != nil {
			return currency.Rate{}, fmt.Errorf("row %d: invalid NumCode %q: %w", row, rawCode, err)
		}

		numCode = parsedCode
	}

	rawVal := strings.TrimSpace(entry.Value)
	if rawVal == "" {
		return currency.Rate{}, fmt.Errorf("row %d: %w", row, errEmptyRateValue)
	}

	normalizedVal := strings.ReplaceAll(rawVal, ",", ".")

	val, err := strconv.ParseFloat(normalizedVal, 64)
	if err != nil {
		return currency.Rate{}, fmt.Errorf("row %d: invalid Value %q: %w", row, rawVal, err)
	}

	if math.IsNaN(val) || math.IsInf(val, 0) {
		return currency.Rate{}, fmt.Errorf("row %d: %w", row, errNonFiniteValue)
	}

	return currency.Rate{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(entry.CharCode),
		Value:    val,
	}, nil
}
