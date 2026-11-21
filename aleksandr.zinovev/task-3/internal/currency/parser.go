package currency

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html/charset"
)

var (
	errInvalidRootTag = errors.New("unexpected XML root element, expected ValCurs")
	errEmptyValutes   = errors.New("XML document contains no Valute entries")
)

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func ParseXML(path string) ([]Valute, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	if curs.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("%w: got %q", errInvalidRootTag, curs.XMLName.Local)
	}

	if len(curs.Valutes) == 0 {
		return nil, errEmptyValutes
	}

	for i := range curs.Valutes {
		curs.Valutes[i].CharCode = strings.TrimSpace(curs.Valutes[i].CharCode)
	}

	return curs.Valutes, nil
}
