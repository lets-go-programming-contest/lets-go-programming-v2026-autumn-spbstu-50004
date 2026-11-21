package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/GenJune222/task-3/internal/currency"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

var errBlankOutputPath = errors.New("output JSON path cannot be blank")

func SaveJSON(path string, rates []currency.Rate) error {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return errBlankOutputPath
	}

	if rates == nil {
		rates = []currency.Rate{}
	}

	for i, r := range rates {
		if math.IsNaN(r.Value) || math.IsInf(r.Value, 0) {
			return fmt.Errorf("rate #%d: non-finite value %v", i+1, r.Value)
		}
	}

	payload, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}
	payload = append(payload, '\n')

	cleanPath := filepath.Clean(trimmed)
	parent := filepath.Dir(cleanPath)

	if err := os.MkdirAll(parent, dirPermissions); err != nil {
		return fmt.Errorf("prepare directory %q: %w", parent, err)
	}
	if err := os.WriteFile(cleanPath, payload, filePermissions); err != nil {
		return fmt.Errorf("write JSON %q: %w", cleanPath, err)
	}
	return nil
}
