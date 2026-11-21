package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GenJune222/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

var (
	errBlankXMLPath   = errors.New("XML file path cannot be blank")
	errInvalidRootTag = errors.New("unexpected XML root element, expected ValCurs")
	errNoValutes      = errors.New("XML document contains no Valute entries")
)

type valCurs struct {
	XMLName xml.Name        `xml:"ValCurs"`
	Valutes []currency.Rate `xml:"Valute"`
}

func ParseXML(path string) ([]currency.Rate, error) {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return nil, errBlankXMLPath
	}
	cleanPath := filepath.Clean(trimmed)

	raw, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", cleanPath, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(raw))
	decoder.CharsetReader = charset.NewReaderLabel

	var doc valCurs
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode XML from %q: %w", cleanPath, err)
	}

	if doc.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("%w: got %q", errInvalidRootTag, doc.XMLName.Local)
	}
	if len(doc.Valutes) == 0 {
		return nil, errNoValutes
	}

	return doc.Valutes, nil
}
