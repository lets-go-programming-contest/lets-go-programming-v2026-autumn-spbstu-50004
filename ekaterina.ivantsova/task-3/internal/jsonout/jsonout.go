package jsonout

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekaterinaivanc/task-3/internal/valute"
)

var errEmptyOutputPath = errors.New("output file path is empty")

const (
	dirPerms  os.FileMode = 0o750
	filePerms os.FileMode = 0o600
)

func Write(path string, items []valute.Item) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return errEmptyOutputPath
	}

	data, err := json.MarshalIndent(items, "", "\t")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	data = append(data, '\n')

	err = createParentDir(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, filePerms)
	if err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}

func createParentDir(path string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, dirPerms)
	if err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}

	return nil
}
