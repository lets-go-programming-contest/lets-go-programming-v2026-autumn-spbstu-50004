package json

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/BinaryKenny/task-3/valute"
	"github.com/BinaryKenny/task-3/xml"
)

const dirPermissions = 0o755

func EncodeToJSON(curs xml.ValCurs, path string) {
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

	slices.SortFunc(curs.Valutes, func(i, j valute.Valute) int {
		return cmp.Compare(j.Value, i.Value)
	})

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "  ")

	err = encoder.Encode(curs.Valutes)
	if err != nil {
		panic(err.Error() + " failed encode")
	}
}
