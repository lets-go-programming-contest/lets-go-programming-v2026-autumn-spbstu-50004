package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SoundwaveAI/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

var (
	errBlankXMLPath   = errors.New("xml file path cannot be blank")
	errInvalidRootTag = errors.New("unexpected root XML tag, expected ValCurs")
	errNoValutesFound = errors.New("XML document contains no Valute entries")
)

type cbrCurs struct {
	XMLName xml.Name        `xml:"ValCurs"`
	Entries []currency.Rate `xml:"Valute"`
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

	return curs.Entries, nil
}
