package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const dirPerm = 0o755

type CommaFloat float64

func (cf *CommaFloat) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var rawValue string
	if err := decoder.DecodeElement(&rawValue, &start); err != nil {
		return fmt.Errorf("failed to decode xml element: %w", err)
	}

	normalized := strings.ReplaceAll(strings.TrimSpace(rawValue), ",", ".")

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float value %q: %w", rawValue, err)
	}

	*cf = CommaFloat(parsed)

	return nil
}

type ValCurs struct {
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

func ParseFromXML(inputPath string) []Currency {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Errorf("failed to open input xml file: %w", err))
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic(fmt.Errorf("failed to close input xml file: %w", closeErr))
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Errorf("failed to decode xml: %w", err))
	}

	return valCurs.Valutes
}

func SortDesc(currencies []Currency) {
	sort.Slice(currencies, func(firstIdx, secondIdx int) bool {
		return currencies[firstIdx].Value > currencies[secondIdx].Value
	})
}

func SaveToJSON(outputPath string, currencies []Currency) {
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, dirPerm); err != nil {
		panic(fmt.Errorf("failed to create output directory: %w", err))
	}

	jsonData, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal json: %w", err))
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Errorf("failed to create output file: %w", err))
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic(fmt.Errorf("failed to close output json file: %w", closeErr))
		}
	}()

	if _, err := file.Write(jsonData); err != nil {
		panic(fmt.Errorf("failed to write json to file: %w", err))
	}
}
