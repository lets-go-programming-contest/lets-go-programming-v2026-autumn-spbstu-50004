package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"github.com/ekaterinaivanc/task-3/internal/valute"
	"golang.org/x/net/html/charset"
)

var (
	errInvalidXMLRoot   = errors.New("invalid XML root element")
	errEmptyXMLDocument = errors.New("XML document does not contain Valute elements")
)

type xmlDocument struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Valutes []valute.Item `xml:"Valute"`
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

	return doc.Valutes, nil
}
