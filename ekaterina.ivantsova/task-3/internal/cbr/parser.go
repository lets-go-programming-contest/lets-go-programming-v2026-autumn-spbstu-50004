package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ekaterinaivanc/task-3/internal/valute"
	"golang.org/x/net/html/charset"
)

var (
	errInvalidXMLRoot       = errors.New("invalid XML root element")
	errEmptyXMLDocument     = errors.New("XML document does not contain Valute elements")
	errEmptyCurrencyValue   = errors.New("currency Value field is empty")
	errNonFiniteCurrencyVal = errors.New("currency Value must be a finite number")
)

type xmlDocument struct {
	XMLName xml.Name  `xml:"ValCurs"`
	Valutes []xmlItem `xml:"Valute"`
}

type xmlItem struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ReadFile(path string) ([]valute.Item, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	items, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse XML file %q: %w", path, err)
	}

	return items, nil
}

func Parse(data []byte) ([]valute.Item, error) {
	var doc xmlDocument

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err := decoder.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	if doc.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("%w: got %q", errInvalidXMLRoot, doc.XMLName.Local)
	}

	if len(doc.Valutes) == 0 {
		return nil, errEmptyXMLDocument
	}

	items := make([]valute.Item, 0, len(doc.Valutes))

	for index, raw := range doc.Valutes {
		item, parseErr := parseItem(raw, index+1)
		if parseErr != nil {
			return nil, parseErr
		}

		items = append(items, item)
	}

	return items, nil
}

func parseItem(raw xmlItem, position int) (valute.Item, error) {
	numCode := 0
	numText := strings.TrimSpace(raw.NumCode)

	if numText != "" {
		parsedNum, err := strconv.Atoi(numText)
		if err != nil {
			return valute.Item{}, fmt.Errorf("valute #%d: invalid NumCode %q: %w", position, raw.NumCode, err)
		}

		numCode = parsedNum
	}

	valText := strings.ReplaceAll(strings.TrimSpace(raw.Value), ",", ".")
	if valText == "" {
		return valute.Item{}, fmt.Errorf("valute #%d: %w", position, errEmptyCurrencyValue)
	}

	valFloat, err := strconv.ParseFloat(valText, 64)
	if err != nil {
		return valute.Item{}, fmt.Errorf("valute #%d: invalid Value %q: %w", position, raw.Value, err)
	}

	if math.IsNaN(valFloat) || math.IsInf(valFloat, 0) {
		return valute.Item{}, fmt.Errorf("valute #%d: %w", position, errNonFiniteCurrencyVal)
	}

	return valute.Item{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(raw.CharCode),
		Value:    valFloat,
	}, nil
}
