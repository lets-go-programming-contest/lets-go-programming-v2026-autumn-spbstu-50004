package json

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

const dirPermissions = 0o755

func EncodeToJSON(curs []JSONValute, path string) {
	direction := filepath.Dir(path)

	err := os.MkdirAll(direction, dirPermissions)
	if err != nil {
		panic(fmt.Errorf("failure in making directory: %w", err))
	}

	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Errorf("failed to create output file: %w", err))
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	slices.SortFunc(curs, func(i, j JSONValute) int {
		return cmp.Compare(j.Value, i.Value)
	})

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "  ")

	err = encoder.Encode(curs)
	if err != nil {
		panic(err.Error() + " failed encode")
	}
}
