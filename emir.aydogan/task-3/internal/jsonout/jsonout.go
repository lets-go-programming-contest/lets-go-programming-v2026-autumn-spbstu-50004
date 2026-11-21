package jsonout

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/emiraydgn/task-3/internal/cbr"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

func Save(path string, currencies []cbr.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
