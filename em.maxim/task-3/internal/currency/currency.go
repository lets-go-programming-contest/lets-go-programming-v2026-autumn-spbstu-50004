package currency

import (
	"bytes"
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

const dirPerm fsFileMode = 0o750

const filePerm = 0o600

type fsFileMode = os.FileMode

type valCurs struct {
	Items []valuteNode `xml:"Valute"`
}

type valuteNode struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(strings.TrimSpace(charset)) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return input, nil
	}
}

func ParseXML(data []byte) ([]Rate, error) {
	var feed valCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	if err := decoder.Decode(&feed); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	rates := make([]Rate, 0, len(feed.Items))

	for _, entry := range feed.Items {
		var numCode int

		rawNumCode := strings.TrimSpace(entry.NumCode)
		if rawNumCode != "" {
			parsedCode, err := strconv.Atoi(rawNumCode)
			if err != nil {
				return nil, fmt.Errorf("parse num code: %w", err)
			}

			numCode = parsedCode
		}

		prepared := strings.ReplaceAll(strings.TrimSpace(entry.Value), ",", ".")

		val, err := strconv.ParseFloat(prepared, 64)
		if err != nil {
			return nil, fmt.Errorf("parse valute value: %w", err)
		}

		rates = append(rates, Rate{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(entry.CharCode),
			Value:    val,
		})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return rates, nil
}

func SaveJSON(outPath string, rates []Rate) error {
	dir := filepath.Dir(outPath)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	encoded, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(outPath, encoded, filePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
