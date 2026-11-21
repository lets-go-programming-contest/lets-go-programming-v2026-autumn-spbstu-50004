package currency

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

func SortByValueDesc(valutes []Valute) {
	slices.SortFunc(valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

func SaveJSON(path string, valutes []Valute) error {
	if valutes == nil {
		valutes = []Valute{}
	}

	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	data = append(data, '\n')

	cleanPath := filepath.Clean(path)
	if err := os.MkdirAll(filepath.Dir(cleanPath), dirPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(cleanPath, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON file %q: %w", cleanPath, err)
	}

	return nil
}
