package valcurs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

func Save(path string, valutes []Valute) error {
	data, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, dirPermissions)
	if err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	err = os.WriteFile(path, data, filePermissions)
	if err != nil {
		return fmt.Errorf("write JSON file: %w", err)
	}

	return nil
}
