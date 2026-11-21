package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SoundwaveAI/task-3/internal/currency"
)

const (
	targetDirMode  = 0o755
	targetFileMode = 0o600
)

var errBlankOutputPath = errors.New("target json path cannot be blank")

func SaveRates(targetPath string, items []currency.Rate) error {
	trimmed := strings.TrimSpace(targetPath)
	if trimmed == "" {
		return errBlankOutputPath
	}

	if items == nil {
		items = []currency.Rate{}
	}

	payload, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal rates to json: %w", err)
	}

	payload = append(payload, '\n')

	cleanTarget := filepath.Clean(trimmed)
	parentDir := filepath.Dir(cleanTarget)

	err = os.MkdirAll(parentDir, targetDirMode)
	if err != nil {
		return fmt.Errorf("prepare target directory %q: %w", parentDir, err)
	}

	err = os.WriteFile(cleanTarget, payload, targetFileMode)
	if err != nil {
		return fmt.Errorf("write json file %q: %w", cleanTarget, err)
	}

	return nil
}
