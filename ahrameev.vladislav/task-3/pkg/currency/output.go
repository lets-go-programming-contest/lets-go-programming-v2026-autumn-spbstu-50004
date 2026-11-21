package currency

import (
	"encoding/json"
	"fmt"
	"os"
)

const filePerm = 0o600

func SaveToJSON(results []CurrencyResult, path string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
