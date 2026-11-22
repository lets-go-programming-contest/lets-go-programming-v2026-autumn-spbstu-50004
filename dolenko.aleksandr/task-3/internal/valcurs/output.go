package valcurs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Save(path string, valutes []Valute) error {
	data, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		return fmt.Errorf("write JSON file: %w", err)
	}

	return nil
}
